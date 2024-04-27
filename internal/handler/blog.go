package handler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Blog struct {
	Logger *zap.Logger
}

func (b *Blog) list(c fiber.Ctx) error {
	b.Logger.Info("blog list page is called!")

	return c.Render("blogs-list", fiber.Map{})
}

func (b *Blog) blog(c fiber.Ctx) error {
	var (
		fullName = c.Get("userFullName")
		role     = c.Get("userRole")
		email    = c.Get("userEmail")
	)
	b.Logger.Info("blog page is called!")
	param := c.Params("blogID")

	err := validation.Validate(param,
		validation.Required, // not empty
		is.Int,
	)
	if err != nil {
		b.Logger.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Render("blog", fiber.Map{
		"blogID":   param,
		"fullName": fullName,
		"email":    email,
		"role":     role,
	})
}

func (b *Blog) Register(g fiber.Router) {
	g.Get("/", b.list)
	g.Get("/:blogID", b.blog)
}
