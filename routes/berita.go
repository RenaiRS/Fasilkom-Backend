package routes

import (
	"redesign_backend/controllers"
	"redesign_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func beritaRoutes(api fiber.Router) {
	berita := api.Group("/berita")
	berita.Post("/", middlewares.Auth, controllers.CreateBerita)
	berita.Get("/", controllers.GetBerita)
	berita.Patch("/:id", middlewares.Auth, controllers.UpdateBerita)
	berita.Delete("/:id", middlewares.Auth, controllers.DeleteBerita)
}
