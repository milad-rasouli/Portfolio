package handler

import "github.com/gofiber/fiber/v3"

func Message(c fiber.Ctx, err error) error {
	return c.JSON(map[string]string{"message": err.Error()})
}
