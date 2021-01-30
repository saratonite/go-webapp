package main

import (
	"webapp/handlers"
	"webapp/handlers/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

func main() {

	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Get("/", handlers.GetIndex)
	app.Get("/api", api.GetIndex)

	app.Listen(":3000")

}
