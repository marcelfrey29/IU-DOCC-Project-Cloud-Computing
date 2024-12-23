package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

var validate *validator.Validate

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
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

		if validationErr := validate.Struct(data); validationErr != nil {
			log.Error("Invalid Request Data.", validationErr.Error())
			return c.Status(400).JSON(map[string]string{"message": "Invalid Request Data."})
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
