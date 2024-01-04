package controllers

import (
	// "net/http"
	// "text/template"

	// "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {


	return c.Render("views/index.html", nil)
}


