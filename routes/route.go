package routes

import (
	"github.com/danendra10/ieee_backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Use(middlewares.IsAuthenticated)
	app.Get("/api/v1/air-polution", controllers.GetAllData)
}
