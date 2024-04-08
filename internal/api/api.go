package api

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type API struct {
	endpoint string
	logger   *zap.Logger
}

func New(e string, l *zap.Logger) (*API, error) {
	return &API{
		endpoint: e,
		logger:   l,
	}, nil
}

func (a *API) HomeHandler(c fiber.Ctx) error {
	a.logger.Info("home page is called!")
	return c.JSON("welcome to the page")
}

func (a *API) PostListHandler(c fiber.Ctx) error {
	a.logger.Info("post list page is called!")

	return c.JSON("welcome to the post list page")
}

func (a *API) PostHandler(c fiber.Ctx) error {
	a.logger.Info("post page is called!")

	return c.JSON("welcome to the post " + c.Params("postID"))
}
