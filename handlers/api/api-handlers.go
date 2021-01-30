package api

import "github.com/gofiber/fiber/v2"

func GetIndex(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{"message": "Hello World"})

}
