package routes

import (
	"github.com/danendra10/ieee_backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New())
	// app.Use(middlewares.IsAuthenticated)
	app.Get("/api/v1/air-polution", controllers.GetAllData)
}
