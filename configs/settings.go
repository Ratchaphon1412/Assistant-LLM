package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func AppSettings(app *fiber.App, cfg *Config) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${locals:requestid} ${ip}:${port} ${time} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CLIENT_URL, // หรือ domain ของ frontend
		AllowCredentials: true,           // สำคัญมาก!
	}))

}
