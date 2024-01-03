package routes

import (
	// "net/http"

	"github.com/Marwahkamilaahmad/go-fiber-first.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(c *fiber.App){
	c.Get("/showUser", controllers.UserControllerShow)
	c.Get("/", controllers.Welcome)
}