package handler

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Home struct {
	Logger *zap.Logger
}

func (h *Home) home(c fiber.Ctx) error {
	h.Logger.Info("home page is called!")
	return c.JSON("howm ")

	// return c.Render("pages/home3", fiber.Map{
	// 	"Note": "Hello from home handler!",
	// }, "layouts/base")
}

func (h *Home) Register(g fiber.Router) {
	g.Get("/", h.home)
}
