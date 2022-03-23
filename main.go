package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fact_usecase "github.com/hmrbcnto/go-leni-api/domain/fact/usecases"
	user_usecase "github.com/hmrbcnto/go-leni-api/domain/user/usecases"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/fact_repository"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/user_repository"
)

// Main function
func main() {

	client, err := db.NewMongoClient()

	if err != nil {
		log.Fatal(err)
	}

	// User Use Cases
	userRepo := user_repository.New(client)
	uc := user_usecase.New(userRepo)

	// Fact Uses Cases
	factRepo := fact_repository.New(client)
	fc := fact_usecase.New(factRepo)

	// New instance of app, fiber is to go what express is to node
	app := fiber.New()

	// Routers/Controllers for Users
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := uc.GetUsers(c)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(users)
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		user, err := uc.CreateUser(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(201).JSON(user)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {

		user, err := uc.GetUserById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(user)

	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		user, err := uc.UpdateUserById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(user)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		result, err := uc.DeleteUserById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		// Not found error
		if result.DeletedCount < 1 {
			return c.Status(404).JSON(err.Error())
		}

		return c.Status(200).JSON("Deleted successfully!")
	})

	// Routers/Controllers for Facts
	app.Get("/facts", func(c *fiber.Ctx) error {
		facts, err := fc.GetAllFacts(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(facts)
	})

	app.Post("/facts", func(c *fiber.Ctx) error {
		fact, err := fc.AddFact(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})

	app.Get("/facts/:id", func(c *fiber.Ctx) error {
		fact, err := fc.GetFactById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})

	app.Patch("/facts/:id", func(c *fiber.Ctx) error {
		fact, err := fc.UpdateFactById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})

	app.Delete("/facts/:id", func(c *fiber.Ctx) error {
		deleteResult, err := fc.DeleteFactById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		if deleteResult.DeletedCount < 1 {
			return c.Status(404).JSON("Fact not found")
		}

		return c.Status(200).JSON("Fact successfully deleted!")
	})

	log.Fatal(app.Listen(":9000"))
}
