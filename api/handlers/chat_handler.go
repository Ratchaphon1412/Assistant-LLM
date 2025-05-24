package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ratchaphon1412/assistant-llm/api/presenters"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/database"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/chat"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/contrib/websocket"

	"github.com/gofiber/fiber/v2"
)

func CreateChat(service chat.Service, cfg *configs.Config) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"
		// Access the Fiber context from the websocket connection
		ctx := context.Background()

		redisChanel := fmt.Sprintf("assistant_llm_%s_chat", c.Locals("email").(string))
		// subscribe redis channel
		pubsub := database.Redis.Rd.Subscribe(ctx, redisChanel)
		defer pubsub.Close()

		// Receive messages from the Redis channel
		go func() {
			ch := pubsub.Channel()
			for msg := range ch {
				if err := c.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
					log.Println("write:", err)
					return
				}
			}
		}()

		// Receive messages from websocket
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			// Parse the message to a map
			var chatRequest presenters.ChatRequest
			if err := json.Unmarshal(msg, &chatRequest); err != nil {
				log.Println("unmarshal:", err)
				break
			}
			chat, err := service.CreateChat(&entities.Chat{
				Prompt: chatRequest.Input,
				Researches: []entities.Research{
					{
						Thumbnails: []entities.Thumbnail{},
					},
				},
			})
			if err != nil {
				log.Println("create chat error:", err)
				c.WriteMessage(websocket.TextMessage, []byte("error"))
				break
			}
			// Trigger the Temporal workflow
			_, err = service.TriggerAIWorkflow(ctx, cfg, chat, redisChanel, chatRequest.Input)
			if err != nil {
				log.Println("trigger workflow error:", err)
				c.WriteMessage(websocket.TextMessage, []byte("error"))
				break
			}

		}

	})
}
