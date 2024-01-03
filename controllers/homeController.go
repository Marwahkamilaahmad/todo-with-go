package controllers

import (
	// "net/http"
	// "text/template"
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	// temp, err := template.ParseFiles("views/home/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	// temp.Execute(w,nil)

	return c.Render("views/home/index.html", nil)
}