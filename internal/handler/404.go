package handler

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type P404 struct {
	Logger *zap.Logger
}

func (p404 *P404) NotFound(c fiber.Ctx) error {
	return c.SendString("not found the route")
}
func (p404 *P404) Register(g fiber.Router) {
	g.Get("/", p404.NotFound)
}
