package fact_usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/fact_repository"
	"github.com/hmrbcnto/go-leni-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type FactUseCase interface {
	GetAllFacts(c *fiber.Ctx) ([]models.Fact, error)
	AddFact(c *fiber.Ctx) (*models.Fact, error)
	GetFactById(c *fiber.Ctx) (*models.Fact, error)
	UpdateFactById(c *fiber.Ctx) (*models.Fact, error)
	DeleteFactById(c *fiber.Ctx) (*mongo.DeleteResult, error)
}

type factUseCase struct {
	factRepo fact_repository.FactRepo
}

func New(factRepo fact_repository.FactRepo) FactUseCase {

	return &factUseCase{
		factRepo: factRepo,
	}
}

// Methods
func (fc *factUseCase) GetAllFacts(c *fiber.Ctx) ([]models.Fact, error) {
	return fc.factRepo.GetFacts(c)
}

func (fc *factUseCase) AddFact(c *fiber.Ctx) (*models.Fact, error) {
	return fc.factRepo.CreateFact(c)
}

func (fc *factUseCase) GetFactById(c *fiber.Ctx) (*models.Fact, error) {
	return fc.factRepo.GetFactById(c)
}

func (fc *factUseCase) UpdateFactById(c *fiber.Ctx) (*models.Fact, error) {
	return fc.factRepo.UpdateFactById(c)
}

func (fc *factUseCase) DeleteFactById(c *fiber.Ctx) (*mongo.DeleteResult, error) {
	return fc.factRepo.DeleteFactById(c)
}
