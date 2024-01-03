package routes

import (
	"github.com/Marwahkamilaahmad/go-fiber-first.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App){
	c.Get("/", controllers.UserControllerShow)
}