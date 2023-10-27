package main

import (
	"dapi/config"
	"dapi/db"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func main() {
	app := fiber.New()
	db.Init()

	config.SetupRoutes(app)
	app.Listen(3000)

	defer db.Close()
}
