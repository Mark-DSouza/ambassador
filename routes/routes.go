package routes

import (
	"ambassador/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/admin/register", controllers.Register)
	app.Post("/api/admin/login", controllers.Login)
	app.Get("/api/admin/user", controllers.User)
}
