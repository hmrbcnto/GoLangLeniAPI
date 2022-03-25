package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db"
	router "github.com/hmrbcnto/go-leni-api/routers"
)

// Main function
func main() {

	client, err := db.NewMongoClient()

	if err != nil {
		log.Fatal(err)
	}

	// New instance of app, fiber is to go what express is to node
	app := fiber.New()

	// Initializing routers
	router.InitializeRouters(app, client)

	log.Fatal(app.Listen(":9000"))
}
