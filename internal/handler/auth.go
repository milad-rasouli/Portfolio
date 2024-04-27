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

const (
	TokenTypeAccess  = 1
	TokenTypeRefresh = 2
)

var WrongPasswordOrEmail = errors.New("password or email is wrong")

type Auth struct {
	Logger       *zap.Logger
	UserStore    store.User
	UserPassword *cipher.UserPassword
	JWTToken     *jwt.JWTToken
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
	a.Logger.Info("user is trying to sign in1")

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
	a.Logger.Info("user is trying to sign in")
	{
		token, err = a.JWTToken.RefreshToken.Create(jwt.JWTUser{
			FullName: UserFromDB.FullName,
			Email:    UserFromDB.Email,
			Role:     "admin", //TODO: implement for diffrent roles
		})
		if err != nil {
			a.Logger.Error("Refresh token failed", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		a.Logger.Info("signed in user", zap.Any("user", UserFromDB), zap.String("token:", token))
		SetTokenCookie(c, token, TokenTypeRefresh)
	}

	{
		token, err = a.JWTToken.AccessToken.Create(jwt.JWTUser{
			FullName: UserFromDB.FullName,
			Email:    UserFromDB.Email,
			Role:     "admin", //TODO: implement for diffrent roles
		}) //TODO: Refactor the return arguments of the jwt package
		if err != nil {
			a.Logger.Error("access token failed", zap.Error(err))
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		a.Logger.Info("New access token", zap.String("token", token))
		SetTokenCookie(c, token, TokenTypeAccess)
	}

	return c.Redirect().To("/")
}

// TODO: stop the update request for ordinal users in the middleware
func (a *Auth) updateToken(c fiber.Ctx) error { //TODO: in frontend side should handel the incoming traffic of this route
	var (
		err                           error
		jwtUser                       jwt.JWTUser
		newRefreshToken, refreshToken string
		requestedToken                = c.Get("JWT-Token")
		newAccessToken                string
		accessToken                   string
		accessJWTUser                 jwt.JWTUser
	)
	{
		refreshToken = c.Cookies("jwt_refresh_token")
		if len(refreshToken) == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	{
		jwtUser, err = a.JWTToken.RefreshToken.VerifyParse(refreshToken)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	switch requestedToken {
	case "access":
		{
			accessToken = c.Cookies("jwt_access_token") //TODO: check what happens if the cookie is removed
			accessJWTUser, err = a.JWTToken.AccessToken.VerifyParse(accessToken)
			if accessJWTUser.Email != "" && accessJWTUser.Email != jwtUser.Email {
				a.Logger.Error("Access Token failed match", zap.String("err", "refresh token email: "+jwtUser.Email+" access token email: "+accessJWTUser.Email))
			} else if err != nil {
				newAccessToken, err = a.JWTToken.AccessToken.Create(jwtUser)
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}
				SetTokenCookie(c, newAccessToken, TokenTypeAccess)
			}
			return c.SendStatus(fiber.StatusOK)
		}
	case "refresh":

		{
			since := time.Since(jwtUser.InitiateTime)
			a.Logger.Info(since.String())
			if since < time.Second*20 { //TODO: change it after the task
				a.Logger.Info("unneccery request")
				return c.SendStatus(fiber.StatusAccepted)
			}
		}
		{
			newRefreshToken, err = a.JWTToken.RefreshToken.Create(jwtUser)
			if err != nil {
				a.Logger.Error("refresh token failed", zap.Error(err))
				return c.SendStatus(fiber.StatusInternalServerError)
			}
			a.Logger.Info("New refresh token", zap.String("token", newRefreshToken))
			SetTokenCookie(c, newRefreshToken, TokenTypeRefresh)
		}
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (a *Auth) Register(g fiber.Router) {
	g.Get("/sign-up", a.GetSignUp)
	g.Post("/sign-up", a.PostSignUp)
	g.Get("/sign-in", a.GetSignIn)
	g.Post("/sign-in", a.PostSignIn)

	g.Post("update-token", a.updateToken)
}

func SetTokenCookie(c fiber.Ctx, token string, token_type int) {
	var (
		expTime time.Time
		path    string
		name    string
	)

	if token_type == TokenTypeRefresh {
		expTime = time.Now().Add(time.Second * jwt.RefreshTokenExpireAfter) //TODO: turn it to Hour
		path = "/user/update-token"
		name = "jwt_refresh_token"
	} else {
		expTime = time.Now().Add(time.Second * jwt.AccessTokenExpireAfter) //TODO: turn it to Min
		path = "/"
		name = "jwt_access_token"
	}
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expTime,
		HTTPOnly: true,
		Secure:   true, // false for when you do not use Https
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     path,
		// Domain:   "MiladRasouli.ir", //TODO: take it from the config
	})
}
