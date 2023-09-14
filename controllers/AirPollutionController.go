package controllers

import (
	"log"

	"github.com/danendra10/ieee_backend/database"
	"github.com/danendra10/ieee_backend/models"
	"github.com/danendra10/ieee_backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllData(c *fiber.Ctx) error {
	var get_data []models.AirPolution

	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: ", r)
		}
	}()

	database.DB.Preload("AirPolution").Find(&get_data)

	return utils.JSONResponse(c, fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    get_data,
	})
}
