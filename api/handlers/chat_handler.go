package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ratchaphon1412/assistant-llm/api/presenters"
	"github.com/Ratchaphon1412/assistant-llm/cmd/driver/temporal"
	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/Ratchaphon1412/assistant-llm/pkg/chat"
	"github.com/Ratchaphon1412/assistant-llm/pkg/entities"
	"github.com/gofiber/contrib/websocket"

	"github.com/gofiber/fiber/v2"
)

func CreateChat(service chat.Service, cfg *configs.Config) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			var chatRequest *presenters.ChatRequest

			err = json.Unmarshal(msg, &chatRequest)
			if err != nil {
				log.Println("json unmarshal error:", err)
				c.WriteMessage(websocket.TextMessage, []byte("error"))
				break

			}

			chat, err := service.CreateChat(&entities.Chat{
				Prompt: chatRequest.Input,
			})
			temporal.AIWorkflow(cfg, chat.WorkflowID, chatRequest.Input)
			if err != nil {
				log.Println("create chat error:", err)
				c.WriteMessage(websocket.TextMessage, []byte("error"))
				break
			}

			err = c.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
}
