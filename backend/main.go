package main

import (
	"errors"

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
			return c.Status(500).JSON(map[string]string{"message": "Error while getting Travel Guides."})
		}
		return c.Status(200).JSON(tgs)
	})

	// Get a Travel Guide
	app.Get("/travel-guides/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		auth := c.Get("x-tg-secret")

		travelGuide, err := getTravelGuide(id, auth)
		var unauthorizedError *UnauthorizedError
		var notFoundError *NotFoundError
		if errors.As(err, &unauthorizedError) {
			log.Error("The Secret is not valid.", err.Error())
			return c.Status(401).JSON(map[string]string{"message": "The Secret is not valid."})
		} else if errors.As(err, &notFoundError) {
			log.Error("Travel Guide with the given ID doesn't exist.", id)
			return c.Status(404).JSON(map[string]string{"message": "Travel Guide doesn't exist."})
		} else if err != nil {
			log.Error("Error while getting a Travel Guide.", err.Error())
			return c.Status(500).JSON(map[string]string{"message": "Error while getting the Travel Guide."})
		}
		return c.Status(200).JSON(travelGuide)
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

	// Delete a Travel Guide
	app.Delete("/travel-guides/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		auth := c.Get("x-tg-secret")

		// Get the Travel Guide and check the secret value
		err := deleteTravelGuide(id, auth)
		var unauthorizedError *UnauthorizedError
		var notFoundError *NotFoundError
		if errors.As(err, &unauthorizedError) {
			log.Error("The Secret is not valid.", err.Error())
			return c.Status(401).JSON(map[string]string{"message": "The Secret is not valid."})
		} else if errors.As(err, &notFoundError) {
			// If the Travel Guide doesn't exist in the Database, it is already deleted => Return success because
			// we're in the expected state.
			log.Info("Travel Guide with the given ID doesn't exist: Already in expected state (Deleted).", id)
			return c.Status(200).JSON(map[string]string{"message": "Success."})
		} else if err != nil {
			log.Error("Error while deleting Travel Guide.", err.Error())
			return c.Status(500).JSON(map[string]string{"message": "Error while deleting the Travel Guide."})
		}

		return c.Status(200).JSON(map[string]string{"message": "Success."})
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
			privateTravelGuide.Id = item.TravelGuide.Id
			travelGuides = append(travelGuides, *privateTravelGuide)
		} else {
			log.Debug("Travel Guide is Public, returning.", item.HashId, item.RangeId)
			travelGuides = append(travelGuides, item.TravelGuide)
		}
	}
	log.Info("Got all Travel Guides.", len(travelGuides))

	return travelGuides, nil
}

// Get a single Travel Guide by ID.
func getTravelGuide(id string, secret string) (TravelGuide, error) {
	item, err := GetTravelGuideFromDDB(id)
	if err != nil {
		log.Error("Error while getting Travel Guide.", error.Error)
		return TravelGuide{}, err
	}

	travelGuide := item.TravelGuide
	log.Debug("Got Travel Guide from DynamoDB, performing additional checks.", travelGuide)
	if travelGuide.Private {
		log.Debug("Travel Guide is  Private, checking Secret.")

		err := bcrypt.CompareHashAndPassword([]byte(item.Secret), []byte(secret))
		if err != nil {
			log.Warn("The provided secret doesn't match the Travel Guide Secret: Unauthorized.")
			return travelGuide, &UnauthorizedError{message: "The secret is not valid."}
		}
	}

	return item.TravelGuide, nil
}

// Create a new Travel Guide.
func createTravelGuide(travelGuide *CreateTravelGuideRequest) (TravelGuide, string, error) {
	id := "TG_" + uuid.NewString()
	secret := getBcrypSecretFromPlaintext(travelGuide.Secret)
	travelGuide.TravelGuide.Id = id
	tgi := TravelGuideItem{
		HashId:      "TG",
		RangeId:     id,
		Secret:      secret,
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

// Delete a Travel Guide by ID.
func deleteTravelGuide(id string, secret string) error {
	// Get Travel Guide
	item, err := GetTravelGuideFromDDB(id)
	if err != nil {
		log.Error("Error while getting Travel Guide.", error.Error)
		return err
	}
	log.Debug("Got Travel Guide from DynamoDB, performing additional checks.", item.HashId, item.RangeId)

	// Check Secret
	err = bcrypt.CompareHashAndPassword([]byte(item.Secret), []byte(secret))
	if err != nil {
		log.Warn("The provided secret doesn't match the Travel Guide Secret: Unauthorized.")
		return &UnauthorizedError{message: "The secret is not valid."}
	}

	// Delete
	err = DeleteGuideFromDDB(id)
	if err != nil {
		log.Error("Error while deleting Travel Guide.", err.Error())
		return err
	}

	return nil
}

func getBcrypSecretFromPlaintext(secret string) string {
	encryptedSecret, _ := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	return string(encryptedSecret)
}
