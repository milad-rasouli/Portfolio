package jwt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenStuff(t *testing.T) {
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
		token, err = jwt.CreateRefreshToken(user)
		assert.NoError(t, err)
		fmt.Printf("created token: %s\n", token)
	}

	{
		jwtUser, err := jwt.VerifyParseRefreshToken(token)
		assert.NoError(t, err)
		fmt.Printf("parsed token: %+v\n", jwtUser)
	}

	{
		expiredToken := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJNaWxhZFJhc291bGkuaXIiLCJzdWIiOiJmb28gYmFyIiwiYXVkIjpbImFkbWluIl0sImV4cCI6MTcxNDAzNDE4MCwibmJmIjoxNzE0MDM0MTMyLCJpYXQiOjE3MTQwMzQxMzIsImp0aSI6ImJhckBiYXouY29tIn0.1sdIsJ1XGm4kCqEha98TvAyoEFqJ74sRcI1zjHN6LFKqd-GfzPpwnL-_9BIPeTC1B43ZrUaDMbHy8KhiOjMcBg"
		jwtUser, err := jwt.VerifyParseRefreshToken(expiredToken)
		assert.Error(t, err)

		fmt.Printf("parsed token: %+v expiration error: %+v\n", jwtUser, err)
	}

}
