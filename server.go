package main

import (
	"webapp/database"
	"webapp/handlers"
	"webapp/handlers/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

func main() {

	// Init database connection

	database.InitDb()

	// Initialize view engine
	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{Views: engine})

	// Serve Static files (css/js/images)
	app.Static("/static", "./static")

	app.Get("/", handlers.GetIndex)
	app.Get("/api/todo", api.GetIndex)
	app.Post("/api/todo", api.CreateTodo)
	app.Delete("/api/todo/:id", api.DeleteTodo)

	//Listen port
	app.Listen(":3000")

}
