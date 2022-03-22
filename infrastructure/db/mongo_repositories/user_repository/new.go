package user_repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetUsers(c *fiber.Ctx) ([]models.User, error)
	CreateUser(c *fiber.Ctx) (*models.User, error)
	GetUserById(c *fiber.Ctx) (*models.User, error)
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

func (uRepo *userRepo) CreateUser(c *fiber.Ctx) (*models.User, error) {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return nil, err
	}

	user.ID = ""

	insertionResult, err := uRepo.db.Collection("users").InsertOne(c.Context(), user)

	if err != nil {
		return nil, err
	}

	/// Querying to make sure data was inserted
	// Building request
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Query
	createdRecord := uRepo.db.Collection("users").FindOne(c.Context(), filter)

	// Decoding
	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	// Returning
	return createdUser, nil
}

func (uRepo *userRepo) GetUserById(c *fiber.Ctx) (*models.User, error) {
	id := c.Params("id")

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "_id", Value: userId}}

	// Query
	foundRecord := uRepo.db.Collection("users").FindOne(c.Context(), query)

	// Decoding
	foundUser := &models.User{}
	foundRecord.Decode(foundUser)

	// Returning
	return foundUser, nil

}
