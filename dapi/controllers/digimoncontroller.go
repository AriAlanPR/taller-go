package controllers

import (
	"dapi/models"

	"github.com/gofiber/fiber"
)

// list all digimons.
func GetDigimons(c *fiber.Ctx) {
	c.JSON(models.GetDigimons())
}

// search for a digimon by id.
func GetDigimon(c *fiber.Ctx) {
	id := c.Params("id")
	c.JSON(models.GetDigimon(id))
}

// create a new digimon with received parameters.
func NewDigimon(c *fiber.Ctx) {
	c.JSON(models.NewDigimon(1, "Alphamon", "Ultimate", "Vaccine", "https://digimon-api.com/images/digimon/w/Alphamon.png"))
}

// delete a digimon using received id.
func DeleteDigimon(c *fiber.Ctx) {
	id := c.Params("id")

	if !models.DeleteDigimon(id) {
		c.Status(500).Send("No Digimon found with that ID")
	}
	c.Send("Digimon Successfully deleted")
}
