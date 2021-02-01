package api

import (
	"webapp/database"

	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx) error {

	var todos []database.Todo

	database.Conn.Find(&todos)

	return c.JSON(todos)

}

func CreateTodo(c *fiber.Ctx) error {

	todo := new(database.Todo)

	if err := c.BodyParser(todo); err != nil {

		return c.Status(503).SendString("Error")

	}

	database.Conn.Create(&todo)

	return c.JSON(todo)

}

func DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	var todo database.Todo

	if err := database.Conn.First(&todo, id).Error; err != nil {
		return c.Status(404).SendString("Todo not found")
	}

	database.Conn.Delete(&todo)

	return c.Status(200).SendString("Deleted")

}
