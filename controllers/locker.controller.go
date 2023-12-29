package controllers

import (
	"gorm-relation-tutorial/config"
	"gorm-relation-tutorial/models"

	"github.com/gofiber/fiber/v2"
)

func LockerGetAll(c *fiber.Ctx) error {

	var lockers []models.Locker

	config.DB.Preload("User").Find(&lockers)

	return c.Status(200).JSON(fiber.Map{
		"message": "lockers data success",
		"data":    lockers,
	})
}

func CreateLocker(c *fiber.Ctx) error {

	locker := new(models.Locker)

	if err := c.BodyParser(locker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "bad request",
		})
	}

	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "code is required",
		})
	}

	if locker.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "user id is required",
		})
	}

	config.DB.Create(&locker)

	return c.Status(201).JSON(fiber.Map{
		"message": "create locker success",
		"data":    locker,
	})

}
