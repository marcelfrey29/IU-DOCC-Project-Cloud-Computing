package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Get Travel Guides
	app.Get("/travel-guides", func(c *fiber.Ctx) error {
		tgs, err := getTravelGuides()
		if err != nil {
			log.Error("Error while getting Travel Guides.", err.Error())
			c.Status(500).JSON(map[string]string{"message": "Error while getting Travel Guides."})
		}
		return c.Status(200).JSON(tgs)
	})

	// Create a new Travel Guide
	app.Post("/travel-guides", func(c *fiber.Ctx) error {
		data := new(CreateTravelGuideRequest)
		err := c.BodyParser(data)
		if err != nil {
			log.Error("Error while parsing body.")
		}

		if validationErr := validate.Struct(data); validationErr != nil {
			log.Error("Invalid Request Data.", validationErr.Error())
			return c.Status(400).JSON(map[string]string{"message": "Invalid Request Data."})
		}

		travelGuide, _, err := createTravelGuide(data)
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

// Get all Travel Guides.
//
// For private Travel Guides, all information (except the name) are removed.
func getTravelGuides() ([]TravelGuide, error) {
	log.Info("Get all Travel Guides.")
	items, err := GetTravelGuidesFromDDB()

	if err != nil {
		log.Error("Error while getting Travel Guides.", error.Error)
		return nil, err
	}

	// Travel Guide Items must be transformed to Travel Guides.
	// In addition, information in private Travel Guides must be removed.
	var travelGuides []TravelGuide
	for _, item := range items {
		if item.TravelGuide.Private {
			log.Debug("Travel Guide is Private, hiding all private information.", item.HashId, item.RangeId)
			privateTravelGuide := new(TravelGuide)
			privateTravelGuide.Private = true
			privateTravelGuide.Name = item.TravelGuide.Name
			travelGuides = append(travelGuides, *privateTravelGuide)
		} else {
			log.Debug("Travel Guide is Public, returning.", item.HashId, item.RangeId)
			travelGuides = append(travelGuides, item.TravelGuide)
		}
	}
	log.Info("Got all Travel Guides.", len(travelGuides))

	return travelGuides, nil
}

// Create a new Travel Guide.
func createTravelGuide(travelGuide *CreateTravelGuideRequest) (TravelGuide, string, error) {
	id := "TG_" + uuid.NewString()
	secret, _ := bcrypt.GenerateFromPassword([]byte(travelGuide.Secret), bcrypt.DefaultCost)
	travelGuide.TravelGuide.Id = id
	tgi := TravelGuideItem{
		HashId:      "TG",
		RangeId:     id,
		Secret:      string(secret),
		TravelGuide: travelGuide.TravelGuide,
	}
	log.Info("Creating a new Travel Guide.", tgi.TravelGuide)

	item, err := CreateTravelGuideInDDB(&tgi)
	if err != nil {
		log.Error("Error while creating Travel Guide.", error.Error)
		return TravelGuide{}, "", err
	}

	tg := item.TravelGuide
	log.Info("Created Travel Guide.", tg)
	return tg, item.Secret, nil
}
