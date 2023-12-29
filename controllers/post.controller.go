package controllers

import (
	"gorm-relation-tutorial/config"
	"gorm-relation-tutorial/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post

	config.DB.Preload("User").Find(&posts)

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func GetPostDetail(c *fiber.Ctx) error {
	var post models.Post

	config.DB.Preload("User").First(&post)

	return c.JSON(fiber.Map{
		"posts": post,
	})
}
