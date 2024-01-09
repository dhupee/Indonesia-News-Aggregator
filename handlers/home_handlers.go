package handlers

import "github.com/gofiber/fiber/v2"

func RootHandler(c *fiber.Ctx) error {
	return c.SendFile("./assets/welcome.txt")
}
