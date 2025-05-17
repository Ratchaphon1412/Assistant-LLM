package middlewares

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func UpgradeRequest(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
		return c.Next()
	}
	return c.SendStatus(fiber.StatusUpgradeRequired)
}
