package routes

import (
	"redesign_backend/controllers"
	"redesign_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func kemahasiswaanKerjasamaRoutes(api fiber.Router) {
	kemahasiswaan := api.Group("/kemahasiswaan-kerjasama")
	kemahasiswaan.Post("/", middlewares.Auth, controllers.CreateKemahasiswaanKerjasama)
	kemahasiswaan.Get("/", controllers.GetKemahasiswaanKerjasama)
	kemahasiswaan.Patch("/:id", middlewares.Auth, controllers.UpdateKemahasiswaanKerjasama)
	kemahasiswaan.Delete("/:id", middlewares.Auth, controllers.DeleteKemahasiswaanKerjasama)
}