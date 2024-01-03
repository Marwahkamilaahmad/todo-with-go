package main

import (

	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	// "github.com/Marwahkamilaahmad/go-fiber-first.git/entity/migration"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	// for migrate
	// migration.RunMigrate()
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message" : "Hello world",
	// 		"data" : "no data shown",
	// 	})
	// })
	routes.RouterApp(app)

	app.Listen(":8000")

	 
}