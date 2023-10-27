package config

import (
	"dapi/controllers"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/digimons", controllers.GetDigimons)
	app.Get("/digimons/:id", controllers.GetDigimon)
	app.Post("/digimons", controllers.NewDigimon)
	app.Delete("/digimons/:id", controllers.DeleteDigimon)
}
