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

func (a *Auth) GetSignUp(c fiber.Ctx) error {
	a.Logger.Info("sign up page is called!")
	return c.Render("sign-up", fiber.Map{})
}

func (a *Auth) PostSignUp(c fiber.Ctx) error {
	// usr := model.UserSignUp{}
	data := c.Body()
	a.Logger.Info(string(data))
	return c.JSON(map[string]string{"message": "just a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple error"})
}

func (a *Auth) GetSignIn(c fiber.Ctx) error {
	a.Logger.Info("sign in page is called!")
	return c.Render("sign-in", fiber.Map{})
}

func (a *Auth) PostSignIn(c fiber.Ctx) error {
	data := c.Body()
	a.Logger.Info(string(data))
	return c.JSON(map[string]string{"message": "just a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple errorjust a simple error"})

}

func (a *Auth) Register(g fiber.Router) {
	g.Get("/sign-up", a.GetSignUp)
	g.Post("/sign-up", a.PostSignUp)
	g.Get("/sign-in", a.GetSignIn)
	g.Post("/sign-in", a.PostSignIn)
}
