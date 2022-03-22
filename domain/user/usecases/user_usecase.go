package user_usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmrbcnto/go-leni-api/infrastructure/db/mongo_repositories/user_repository"
	"github.com/hmrbcnto/go-leni-api/models"
)

type UserUseCase interface {
	GetUsers(c *fiber.Ctx) ([]models.User, error)
	CreateUser(c *fiber.Ctx) (*models.User, error)
	GetUserById(c *fiber.Ctx) (*models.User, error)
	UpdateUserById(c *fiber.Ctx) (*models.User, error)
}

type userUseCase struct {
	userRepo user_repository.UserRepo
}

func New(userRepo user_repository.UserRepo) UserUseCase {

	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) GetUsers(c *fiber.Ctx) ([]models.User, error) {
	return uc.userRepo.GetUsers(c)
}

func (uc *userUseCase) CreateUser(c *fiber.Ctx) (*models.User, error) {
	return uc.userRepo.CreateUser(c)
}

func (uc *userUseCase) GetUserById(c *fiber.Ctx) (*models.User, error) {
	return uc.userRepo.GetUserById((c))
}

func (uc *userUseCase) UpdateUserById(c *fiber.Ctx) (*models.User, error) {
	return uc.userRepo.UpdateUserById(c)
}
