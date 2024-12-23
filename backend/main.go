package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Create a new Travel Guide
	app.Post("/travel-guides", func(c *fiber.Ctx) error {
		data := new(TravelGuide)
		err := c.BodyParser(data)
		if err != nil {
			log.Error("Error while parsing body.")
		}
		data.Id = uuid.NewString()
		log.Info("Creating a new Travel Guide.", data)

		travelGuide, err := CreateTravelGuide(data)
		if err != nil {
			return c.Status(500).JSON(map[string]string{"message": "Error while creating Travel Guide."})
		}
		log.Info("Created Travel Guide.", travelGuide)
		return c.Status(201).JSON(travelGuide)
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
	})

	app.Listen(":3000")
}
