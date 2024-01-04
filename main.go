package main

import (
	"net/http"

	"github.com/Marwahkamilaahmad/go-fiber-first.git/controllers"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity/migration"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/routes"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	migration.RunMigrate()

	app := fiber.New()
	app.Get("/", controllers.Welcome)

	
	routes.RouterApp()

	http.ListenAndServe(":8000", nil)
	// app.Listen(":8000")

	 
}