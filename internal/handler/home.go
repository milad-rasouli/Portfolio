package handler

import (
	"errors"

	"github.com/Milad75Rasouli/portfolio/frontend/views/pages"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Home struct {
	Logger    *zap.Logger
	HomeStore store.Home
}

func (h *Home) home(c fiber.Ctx) error {
	var (
		home model.Home
		err  error
	)
	{
		home, err = h.HomeStore.GetHome(c.Context())
		if errors.Is(err, store.HomeNotFountError) == false && err != nil {
			h.Logger.Error("home error", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	base := pages.Home(home)
	base.Render(c.Context(), c.Response().BodyWriter())
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendStatus(fiber.StatusOK)
	// return c.Render("home", fiber.Map{
	// 	"name":       home.Name,
	// 	"slogan":     home.Slogan,
	// 	"shortIntro": home.ShortIntro,
	// 	"githubUrl":  home.GithubUrl,
	// })
}

func (h *Home) Register(g fiber.Router) {
	g.Get("/", h.home)
}
