package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRefreshTokenStuff(t *testing.T) {
	var (
		token string
		err   error
	)
	cfg := Config{
		RefreshSecretKey: "VerySecretKey",
	}
	user := JWTUser{
		FullName: "foo bar",
		Email:    "bar@baz.com",
		Role:     "admin",
	}
	jwt := NewRefreshJWT(cfg)

	{
		token, err = jwt.Create(user)
		assert.NoError(t, err)
		fmt.Printf("created token: %s\n", token)
	}

	{
		jwtUser, err := jwt.VerifyParse(token)
		assert.NoError(t, err)
		fmt.Printf("parsed token: %+v\n", jwtUser)
	}

	{
		expiredToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJNaWxhZFJhc291bGkuaXIiLCJzdWIiOiJmb28gYmFyIiwiYXVkIjpbImFkbWluIl0sImV4cCI6MTcxNDAzNDE4MCwibmJmIjoxNzE0MDM0MTMyLCJpYXQiOjE3MTQwMzQxMzIsImp0aSI6ImJhckBiYXouY29tIn0.1sdIsJ1XGm4kCqEha98TvAyoEFqJ74sRcI1zjHN6LFKqd-GfzPpwnL-_9BIPeTC1B43ZrUaDMbHy8KhiOjMcBg"
		jwtUser, err := jwt.VerifyParse(expiredToken)
		assert.Error(t, err)
		fmt.Printf("parsed token: %+v expiration error: %+v\n", jwtUser, err)
	}

}

func TestAccessTokenStuff(t *testing.T) {
	var (
		token string
		err   error
	)
	cfg := Config{
		AccessSecretKey: "VerySecretKey",
	}
	user := JWTUser{
		FullName: "foo bar",
		Email:    "bar@test.com",
		Role:     "admin",
	}
	jwt := NewAccessJWT(cfg)

	{
		token, err = jwt.Create(user)
		assert.NoError(t, err)
		fmt.Printf("created token: %s\n", token)
	}

	{
		jwtUser, err := jwt.VerifyParse(token)
		if err != nil {
			assert.NoError(t, err)
		}
		fmt.Printf("parsed token: %+v\n", jwtUser)

	}
	{
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJhckBiYXouY29tIiwiZXhwIjoxNzE0MDc4MTYxLCJmdWxsX25hbWUiOiJmb28gYmFyIiwicm9sZSI6ImFkbWluIn0.pLMe3zPYggvKpA8SDL2mbV6kfhISW1IH-xgBClS40TI"
		jwtUser, err := jwt.VerifyParse(expiredToken)
		assert.Error(t, err)
		fmt.Printf("parsed token: %+v expiration error: %+v\n", jwtUser, err)
	}
}
