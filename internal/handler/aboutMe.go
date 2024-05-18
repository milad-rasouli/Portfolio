package handler

import (
	"html/template"

	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AboutMe struct {
	Logger       *zap.Logger
	AboutMeStore store.AboutMe
}

func (am *AboutMe) GetAboutMe(c fiber.Ctx) error {
	aboutMe, err := am.AboutMeStore.GetAboutMe(c.Context())
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Render("about-me", fiber.Map{"content": template.HTML(aboutMe.Content)})
}

func (am *AboutMe) Register(g fiber.Router) {
	g.Get("/", am.GetAboutMe)
}