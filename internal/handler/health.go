package handler

import "github.com/gofiber/fiber/v3"

func GetHealth(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
