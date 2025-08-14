package routes

import "github.com/gofiber/fiber/v2"

func RoutesList(app *fiber.App) {
	api := app.Group("/api")
	adminRoutes(api)
	beritaRoutes(api)
	fileRoutes(api)
	kemahasiswaanKerjasamaRoutes(api)
	scholarshipRoutes(api)
	vacanciesRoutes(api)
}