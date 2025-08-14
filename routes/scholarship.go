package routes

import (
	"redesign_backend/controllers"
	"redesign_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func scholarshipRoutes(api fiber.Router) {
	scholarship := api.Group("/scholarship")
	scholarship.Post("/", middlewares.Auth, controllers.CreateScholarship)
	scholarship.Get("/", controllers.GetScholarship)
	scholarship.Patch("/:id", middlewares.Auth, controllers.UpdateScholarship)
	scholarship.Delete("/:id", middlewares.Auth, controllers.DeleteScholarship)
}