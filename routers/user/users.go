package user_router

import (
	"github.com/gofiber/fiber/v2"
	user_usecase "github.com/hmrbcnto/go-leni-api/domain/user/usecases"
)

func RouterInit(app *fiber.App, uc user_usecase.UserUseCase) {
	getAllUsersRoute(app, uc)
	createUserRoute(app, uc)
	getUserByIdRoute(app, uc)
	updateUserByIdRoute(app, uc)
	deleteUserRoute(app, uc)
}

func getAllUsersRoute(app *fiber.App, uc user_usecase.UserUseCase) {
	app.Get("/users", func(c *fiber.Ctx) error {
		users, err := uc.GetUsers(c)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(users)
	})
}

func createUserRoute(app *fiber.App, uc user_usecase.UserUseCase) {
	app.Post("/users", func(c *fiber.Ctx) error {

		user, err := uc.CreateUser(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(201).JSON(user)
	})
}

func getUserByIdRoute(app *fiber.App, uc user_usecase.UserUseCase) {
	app.Get("/users/:id", func(c *fiber.Ctx) error {

		user, err := uc.GetUserById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(user)

	})
}

func updateUserByIdRoute(app *fiber.App, uc user_usecase.UserUseCase) {
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		user, err := uc.UpdateUserById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(user)
	})
}

func deleteUserRoute(app *fiber.App, uc user_usecase.UserUseCase) {
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
}
