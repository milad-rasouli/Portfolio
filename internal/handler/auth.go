package handler

import (

	// validation "github.com/go-ozzo/ozzo-validation"
	// "github.com/go-ozzo/ozzo-validation/is"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Auth struct {
	Logger *zap.Logger
}

func (a *Auth) SignUp(c fiber.Ctx) error {
	a.Logger.Info("sign up page is called!")
	return c.Render("sign-up", fiber.Map{})
}

func (a *Auth) SignIn(c fiber.Ctx) error {

	return c.JSON("SignIn")
}

func (a *Auth) Register(g fiber.Router) {
	g.Get("/sign-up", a.SignUp)
	g.Get("/sign-in", a.SignIn)
}
