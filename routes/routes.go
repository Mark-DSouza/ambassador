package routes

import (
	"ambassador/controllers"
	"ambassador/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
	app.Post("/api/admin/login", controllers.Login)

	authMiddleware := app.Use(middlewares.IsAuthenticated)
	authMiddleware.Get("/api/admin/user", controllers.User)
	authMiddleware.Post("/api/admin/logout", controllers.Logout)
	authMiddleware.Put("/api/admin/user/info", controllers.UpdateInfo)
	authMiddleware.Put("/api/admin/user/password", controllers.UpdatePassword)
}
