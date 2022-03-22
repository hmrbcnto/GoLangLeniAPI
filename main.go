package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	user_usecase "github.com/hmrbcnto/go-leni-api/domain/user/usecases"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/user_repository"
)

// Main function
func main() {

	client, err := db.NewMongoClient()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := user_repository.New(client)
	uc := user_usecase.New(userRepo)

	// New instance of app, fiber is to go what express is to node
	app := fiber.New()

	// Routers/Controllers
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

	// app.Delete("/users/:id", func(c *fiber.Ctx) error {
	// 	idParam := c.Params("id")

	// 	userId, err := primitive.ObjectIDFromHex(idParam)

	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	query := bson.D{{Key: "_id", Value: userId}}

	// 	result, err := mg.Db.Collection("users").DeleteOne(c.Context(), query)

	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	if result.DeletedCount < 1 {
	// 		return c.SendStatus(404)
	// 	}

	// 	return c.Status(200).JSON("Record deleted")
	// })

	log.Fatal(app.Listen(":9000"))
}
