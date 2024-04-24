package handler

import (
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
	var user request.UserSingUp
	c.Bind().Body(&user)
	err := user.Validate()
	if err != nil {
		return Message(c, err) // TODO: retrieve meaningful message based on the error
	}
	now := time.Now()
	validUser := model.User{
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  a.UserPassword.HashPassword(user.Password, user.Email),
		OnlineAt:  now,
		CreatedAt: now,
	}
	validUser.ID, err = a.UserStore.Create(c.Context(), validUser)
	if errors.Is(err, store.DuplicateUserError) {
		return Message(c, errors.New("user is duplicated"))
	} else if err != nil {
		a.Logger.Error("creating user failed", zap.Error(err))
		return Message(c, errors.New("unknown error"))
	}
	a.Logger.Info("user created", zap.Any("User", validUser))
	return c.Redirect().To("/user/sign-in")
}

func (a *Auth) GetSignIn(c fiber.Ctx) error {
	a.Logger.Info("sign in page is called!")
	return c.Render("sign-in", fiber.Map{})
}

func (a *Auth) PostSignIn(c fiber.Ctx) error {
	var user request.UserSignIn
	c.Bind().Body(&user)
	err := user.Validate()
	if err != nil {
		return Message(c, err)
	}

	UserFromDB, err := a.UserStore.GetByEmail(c.Context(), user.Email)
	if err != nil {
		return Message(c, errors.New("password or email is wrong"))
	}
	if a.UserPassword.ComparePasswords(UserFromDB.Password, user.Password, user.Email) == false {
		return Message(c, errors.New("password or email is wrong"))
	}

	a.Logger.Info("signed up user", zap.Any("user", UserFromDB))
	// a.Logger.Info("token", zap.Any("user token", tokenString))
	return Message(c, errors.New("you are signed up")) //TODO: redirect if necessary
}

func (a *Auth) Register(g fiber.Router) {
	g.Get("/sign-up", a.GetSignUp)
	g.Post("/sign-up", a.PostSignUp)
	g.Get("/sign-in", a.GetSignIn)
	g.Post("/sign-in", a.PostSignIn)
}
