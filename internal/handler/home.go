package handler

import (
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
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	return c.Render("home", fiber.Map{
		"name":       home.Name,
		"slogan":     home.Slogan,
		"shortIntro": home.ShortIntro,
		"githubUrl":  home.GithubUrl,
	})
}

func (h *Home) Register(g fiber.Router) {
	g.Get("/", h.home)
}
