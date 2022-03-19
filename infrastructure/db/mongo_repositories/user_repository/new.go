package user_repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetUsers(c *fiber.Ctx) ([]models.User, error)
}

type userRepo struct {
	db *mongo.Database
}

func New(dbClient *mongo.Client) UserRepo {
	return &userRepo{
		db: dbClient.Database("leniApi"),
	}
}

func (uRepo *userRepo) GetUsers(c *fiber.Ctx) ([]models.User, error) {
	query := bson.D{{}}

	cursor, err := uRepo.db.Collection("users").Find(c.Context(), query)

	var users []models.User = make([]models.User, 0)
	if err = cursor.All(c.Context(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
