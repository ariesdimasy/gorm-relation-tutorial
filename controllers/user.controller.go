package controllers

import (
	"gorm-relation-tutorial/config"
	"gorm-relation-tutorial/models"

	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User

	config.DB.Preload("Locker").Find(&users)

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"err": "bad request",
		})
	}

	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	config.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "create data sucess",
		"data":    user,
	})
}
