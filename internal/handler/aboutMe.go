package handler

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AboutMe struct {
	Logger *zap.Logger
}

func (am *AboutMe) GetAboutMe(c fiber.Ctx) error {
	return c.Render("about-me", fiber.Map{"content": "there is nothing to show you bro"})
}

func (am *AboutMe) Register(g fiber.Router) {
	g.Get("/", am.GetAboutMe)
}
