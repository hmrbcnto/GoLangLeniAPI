package fact_router

import (
	"github.com/gofiber/fiber/v2"
	fact_usecase "github.com/hmrbcnto/go-leni-api/domain/fact/usecases"
)

func RouterInit(app *fiber.App, fc fact_usecase.FactUseCase) {
	getAllFactsRoute(app, fc)
	createFactRoute(app, fc)
	getFactByIdRoute(app, fc)
	updateFactById(app, fc)
	deleteFactById(app, fc)
}

func getAllFactsRoute(app *fiber.App, fc fact_usecase.FactUseCase) {
	app.Get("/facts", func(c *fiber.Ctx) error {
		facts, err := fc.GetAllFacts(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(facts)
	})
}

func createFactRoute(app *fiber.App, fc fact_usecase.FactUseCase) {
	app.Post("/facts", func(c *fiber.Ctx) error {
		fact, err := fc.AddFact(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})
}

func getFactByIdRoute(app *fiber.App, fc fact_usecase.FactUseCase) {
	app.Get("/facts/:id", func(c *fiber.Ctx) error {
		fact, err := fc.GetFactById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})
}

func updateFactById(app *fiber.App, fc fact_usecase.FactUseCase) {
	app.Patch("/facts/:id", func(c *fiber.Ctx) error {
		fact, err := fc.UpdateFactById(c)

		if err != nil {
			return c.Status(500).JSON(err.Error())
		}

		return c.Status(200).JSON(fact)
	})
}

func deleteFactById(app *fiber.App, fc fact_usecase.FactUseCase) {
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

}
