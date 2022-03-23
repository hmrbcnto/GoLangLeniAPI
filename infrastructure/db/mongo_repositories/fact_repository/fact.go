package fact_repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FactRepo interface {
	GetFacts(c *fiber.Ctx) ([]models.Fact, error)
	CreateFact(c *fiber.Ctx) (*models.Fact, error)
	GetFactById(c *fiber.Ctx) (*models.Fact, error)
	UpdateFactById(c *fiber.Ctx) (*models.Fact, error)
	DeleteFactById(c *fiber.Ctx) (*mongo.DeleteResult, error)
}

type factRepo struct {
	db *mongo.Database
}

func New(dbClient *mongo.Client) FactRepo {
	return &factRepo{
		db: dbClient.Database("leniApi"),
	}
}

func (fRepo *factRepo) GetFacts(c *fiber.Ctx) ([]models.Fact, error) {
	// Building query
	query := bson.D{{}}

	// Getting query output
	cursor, err := fRepo.db.Collection("facts").Find(c.Context(), query)

	if err != nil {
		return nil, err
	}

	var facts []models.Fact = make([]models.Fact, 0)
	err = cursor.All(c.Context(), &facts)

	if err != nil {
		return nil, err
	}

	return facts, nil
}

func (fRepo *factRepo) CreateFact(c *fiber.Ctx) (*models.Fact, error) {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return nil, err
	}

	fact.ID = ""

	insertionResult, err := fRepo.db.Collection("facts").InsertOne(c.Context(), fact)

	if err != nil {
		return nil, err
	}

	//Query to make sure result was inserted
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Query
	createdRecord := fRepo.db.Collection("facts").FindOne(c.Context(), filter)

	// Cast to fact model
	createdFact := &models.Fact{}
	createdRecord.Decode(createdFact)

	return createdFact, nil
}

func (fRepo *factRepo) GetFactById(c *fiber.Ctx) (*models.Fact, error) {
	// Getting id
	idParam := c.Params("id")

	// Convert to ObjectId
	factId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return nil, err
	}

	// Building query
	query := bson.D{{Key: "_id", Value: factId}}

	// Querying
	foundRecord := fRepo.db.Collection("facts").FindOne(c.Context(), query)

	// Casting to Fact struct
	foundFact := &models.Fact{}
	foundRecord.Decode(foundFact)

	return foundFact, nil
}

func (fRepo *factRepo) UpdateFactById(c *fiber.Ctx) (*models.Fact, error) {
	// Getting id
	idParam := c.Params("id")

	// Converting to ObjectId
	factId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return nil, err
	}

	// Find query
	query := bson.D{{Key: "_id", Value: factId}}

	/// Getting requestbody
	// Creating struct where body will be housed
	fact := new(models.Fact)

	// Parsing
	err = c.BodyParser(fact)

	if err != nil {
		return nil, err
	}

	// Updating!
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "fact", Value: fact.Fact},
				{Key: "source", Value: fact.Source},
			},
		},
	}

	err = fRepo.db.Collection("facts").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		return nil, err
	}

	// Returning input
	fact.ID = idParam

	return fact, nil
}

func (fRepo *factRepo) DeleteFactById(c *fiber.Ctx) (*mongo.DeleteResult, error) {
	idParam := c.Params("id")

	factId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return nil, err
	}

	// Query
	query := bson.D{{Key: "_id", Value: factId}}

	// Deleting
	result, err := fRepo.db.Collection("facts").DeleteOne(c.Context(), query)

	if err != nil {
		return nil, err
	}

	return result, nil
}
