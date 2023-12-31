package controllers

import (
	"fmt"
	"log"

	"github.com/danendra10/ieee_backend/database"
	"github.com/danendra10/ieee_backend/models"
	"github.com/danendra10/ieee_backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllData(c *fiber.Ctx) error {
	var get_data []models.AirPolution
	request_compound := c.Get("type")

	fmt.Println(request_compound)

	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR: ", r)
		}
	}()

	db := database.DB

	// find data from air_pollutions table using where compound = request_compound
	err := db.Where("compound = ?", request_compound).Select("created_at", "lat", "lng").Find(&get_data).Error

	if err != nil {
		return nil
	}

	return utils.JSONResponse(c, fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    get_data,
	})
}

func CreateData(c *fiber.Ctx) error {
	air_pollution_new_data := new(models.AirPolution)

	if err := c.BodyParser(air_pollution_new_data); err != nil {
		return err
	}

	db := database.DB

	if err := db.Create(air_pollution_new_data).Error; err != nil {
		return utils.JSONResponse(c, fiber.Map{
			"code":    400,
			"message": "Success",
			"data":    err,
		})
	}

	return utils.JSONResponse(c, fiber.Map{
		"code":    200,
		"message": "Success",
		"data":    air_pollution_new_data,
	})
}
