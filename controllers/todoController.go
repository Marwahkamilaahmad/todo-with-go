package controllers

import (
	"github.com/Marwahkamilaahmad/go-fiber-first.git/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	todos := models.GetAllTodos()
	data := map[string]any{
		"todo" : todos,
	}

	return c.Render("views/home/index.html", data)
}

func Add(c *fiber.Ctx) error {

}

func Edit(c *fiber.Ctx) error {

}

func Delete(c *fiber.Ctx) error {

}