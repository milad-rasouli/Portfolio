package handler

import (
	"errors"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/jwt"
	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/Milad75Rasouli/portfolio/internal/request"
	"github.com/Milad75Rasouli/portfolio/internal/store"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

var WrongPasswordOrEmail = errors.New("password or email is wrong")

type Auth struct {
	Logger       *zap.Logger
	UserStore    store.User
	UserPassword *cipher.UserPassword
	RefreshJWT   *jwt.RefreshJWT
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
	var (
		user       request.UserSignIn
		token      string
		UserFromDB model.User
		err        error
	)

	{
		c.Bind().Body(&user)
		err = user.Validate()
		if err != nil {
			return Message(c, err)
		}
	}

	{
		UserFromDB, err = a.UserStore.GetByEmail(c.Context(), user.Email)
		if err != nil {
			return Message(c, WrongPasswordOrEmail)
		}
		if a.UserPassword.ComparePasswords(UserFromDB.Password, user.Password, user.Email) == false {
			return Message(c, WrongPasswordOrEmail)
		}
	}

	{
		token, err = a.RefreshJWT.CreateRefreshToken(jwt.JWTUser{
			FullName: UserFromDB.FullName,
			Email:    UserFromDB.Email,
			Role:     "admin", //TODO: implement for diffrent roles
		})
		if err != nil {
			a.Logger.Error("Refresh token failed", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		a.Logger.Info("signed in user", zap.Any("user", UserFromDB), zap.String("token:", token))

		c.Cookie(&fiber.Cookie{
			Name:     "jwt_token",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * jwt.RefreshTokenExpireAfter),
			HTTPOnly: true,
			Secure:   true, // false for when you do not use Https
			SameSite: fiber.CookieSameSiteStrictMode,
			Path:     "/user/refresh-token",
			// Domain:   "MiladRasouli.ir", //TODO: take it from the config
		})
	}
	return c.Redirect().To("/")
}

func (a *Auth) Register(g fiber.Router) {
	g.Get("/sign-up", a.GetSignUp)
	g.Post("/sign-up", a.PostSignUp)
	g.Get("/sign-in", a.GetSignIn)
	g.Post("/sign-in", a.PostSignIn)
}
