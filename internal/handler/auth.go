package handler

import (

	// validation "github.com/go-ozzo/ozzo-validation"
	// "github.com/go-ozzo/ozzo-validation/is"

	"errors"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/request"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type Auth struct {
	Logger       *zap.Logger
	UserStore    store.User
	UserPassword *cipher.UserPassword
}

func (a *Auth) GetSignUp(c fiber.Ctx) error {
	a.Logger.Info("sign up page is called!")
	return c.Render("sign-up", fiber.Map{})
}

func (a *Auth) PostSignUp(c fiber.Ctx) error {
	// data := c.Body()
	// a.Logger.Info(fmt.Sprintf("%+v", user))
	// a.Logger.Info(string(data))
	var user request.UserSingUp
	c.Bind().Body(&user)
	err := user.Validate()
	if err != nil {
		return c.JSON(map[string]string{"message": err.Error()}) // TODO: retrieve meaningful message based on the error
	}
	validUser := model.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: a.UserPassword.HashPassword(user.Password, user.Email),
		OnlineAt: time.Now(),
	}
	err = a.UserStore.Create(c.Context(), validUser)
	if errors.Is(err, store.DuplicateUserError) {
		return c.JSON(map[string]string{"message": "user is duplicated"})
	} else if err != nil {
		a.Logger.Error("creating user failed", zap.Error(err))
		return c.JSON(map[string]string{"message": "unknown error"})
	}
	a.Logger.Info("user created", zap.Any("User", user))
	return c.Redirect().To("/user/sign-in")
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
