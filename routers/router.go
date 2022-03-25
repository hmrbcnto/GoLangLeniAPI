package router

import (
	"github.com/gofiber/fiber/v2"
	fact_usecase "github.com/hmrbcnto/go-leni-api/domain/fact/usecases"
	user_usecase "github.com/hmrbcnto/go-leni-api/domain/user/usecases"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/fact_repository"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/user_repository"
	fact_router "github.com/hmrbcnto/go-leni-api/routers/fact"
	user_router "github.com/hmrbcnto/go-leni-api/routers/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRouters(app *fiber.App, client *mongo.Client) {
	/// Initializing use cases
	initializeUserRoutes(app, client)
	initializeFactRoutes(app, client)
}

func initializeUserRoutes(app *fiber.App, client *mongo.Client) {
	// User usecase
	userRepo := user_repository.New(client)
	uc := user_usecase.New(userRepo)

	user_router.RouterInit(app, uc)
}

func initializeFactRoutes(app *fiber.App, client *mongo.Client) {
	// Fact usecase
	factRepo := fact_repository.New(client)
	fc := fact_usecase.New(factRepo)

	fact_router.RouterInit(app, fc)
}
