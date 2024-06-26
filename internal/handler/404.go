package handler

import (
	"github.com/Milad75Rasouli/portfolio/frontend/views/pages"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type P404 struct {
	Logger *zap.Logger
}

func (p404 *P404) Middleware(c fiber.Ctx) error {
	base := pages.P404()
	base.Render(c.Context(), c.Response().BodyWriter())
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendStatus(fiber.StatusOK)
}
