package handler

import (
	"context"
	"errors"

	"github.com/Milad75Rasouli/portfolio/frontend/views/pages"
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
	if errors.Is(err, store.AboutMeNotFountError) == false && err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}
	base := pages.AboutMe(aboutMe.Content)
	base.Render(context.Background(), c.Response().BodyWriter())
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendStatus(fiber.StatusOK)
	//return c.Render("about-me", fiber.Map{"content": template.HTML(aboutMe.Content)})
}

func (am *AboutMe) Register(g fiber.Router) {
	g.Get("/", am.GetAboutMe)
}
