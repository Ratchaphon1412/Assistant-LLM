package middlewares

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/Ratchaphon1412/assistant-llm/configs"
	"github.com/gofiber/fiber/v2"
)

func LoggerRequest(cfg *configs.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Log the request method and path
		log.Info("Request received")
		log.Info()
		log.Info("Request Method: ", c.Method(), " Request Path:", c.Path())
		log.Info("Request Query: ", c.Queries())

		// Log the request headers
		for key, values := range c.GetReqHeaders() {
			for _, value := range values {
				log.Info("Header:", key, " ", value)
			}
		}

		// Log the request body if it's not empty
		if c.Body() != nil {
			log.Info("Body:", c.Body())
		} else {
			log.Info("Body: empty")
		}
		// Log Cookie if it exists accessToken
		cookie := c.Cookies(cfg.JWT_COOKIE_NAME)
		if cookie != "" {
			log.Info("Cookie:", cfg.JWT_COOKIE_NAME, " ", cookie)
		} else {
			log.Info("Cookie: not found ", cfg.JWT_COOKIE_NAME)
		}
		log.Info()

		// Continue to the next middleware or handler
		return c.Next()
	}

}
