package main

import (
	"shuttle/databases"
	"shuttle/routes"
	"shuttle/utils"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	
	app.Get("/ws/:id", websocket.New(utils.HandleWebSocketConnection))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} [${status}] ${latency}\n",
	}))

	routes.Route(app)
	database.MongoConnection()

	if err := app.Listen("192.168.110.84:8080"); err != nil {
        panic(err)
    }
	// if err := app.Listen(":8080"); err != nil {
    //     panic(err)
    // }
}