package routes

import (
	"redesign_backend/controllers"
	"redesign_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func vacanciesRoutes(api fiber.Router) {
	vacancies := api.Group("/vacancies")
	vacancies.Post("/", middlewares.Auth, controllers.CreateVacancies)
	vacancies.Get("/", controllers.GetVacancies)
	vacancies.Patch("/:id", middlewares.Auth, controllers.UpdateVacancies)
	vacancies.Delete("/:id", middlewares.Auth, controllers.DeleteVacancies)
}