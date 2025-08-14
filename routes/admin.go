package routes

import (
	"redesign_backend/controllers"
	"redesign_backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func adminRoutes(api fiber.Router) {
	admin := api.Group("/admin")
	admin.Post("/signup", controllers.SignUp)
	admin.Post("/login", controllers.Login)
	admin.Post("/sync-berita", middlewares.Auth, controllers.SyncBeritaCache)
	admin.Post("/logout", middlewares.Auth, controllers.Logout)
}
