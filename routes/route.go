package routes

import (
	"github.com/danendra10/ieee_backend/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
	// 	AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
	// 	ExposeHeaders:    "Origin",
	// 	AllowCredentials: true,
	// }))
	app.Use(cors.New())
	// app.Use(middlewares.IsAuthenticated)
	app.Get("/api/v1/air-polution", controllers.GetAllData)
	app.Post("/api/v1/air-polution/update", controllers.CreateData)
}
