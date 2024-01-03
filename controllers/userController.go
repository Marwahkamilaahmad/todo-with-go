package controllers

import "github.com/gofiber/fiber/v2"


func UserControllerShow(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message" : "hellow world",
	})
}