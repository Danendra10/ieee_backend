package utils

import (
	"github.com/gofiber/fiber/v2"
)

func JSONResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}
