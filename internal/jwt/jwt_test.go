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
		SecretKey: "VerySecretKey",
	}
	user := JWTUser{
		FullName: "foo bar",
		Email:    "bar@baz.com",
		Role:     "admin",
	}
	jwt := NewJWT(cfg)

	{
		token, err = jwt.CreateToken(user)
		assert.NoError(t, err)
		fmt.Printf("created token: %s\n", token)
	}

	{
		jwtUser, err := jwt.VerifyParseUserToken(token)
		assert.NoError(t, err)
		fmt.Printf("parsed token: %+v\n", jwtUser)
	}

	{
		expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJNaWxhZFJhc291bGkuaXIiLCJzdWIiOiJmb28gYmFyIiwiYXVkIjpbImFkbWluIl0sImV4cCI6MTcxNDAzMTM5NCwibmJmIjoxNzE0MDMxMzkzLCJpYXQiOjE3MTQwMzEzOTMsImp0aSI6ImJhckBiYXouY29tIn0.7UVbHcnh3w70TZAywpexFhc3az_S77hYHcUezM19xtM"
		jwtUser, err := jwt.VerifyParseUserToken(expiredToken)
		assert.Error(t, err)

		fmt.Printf("parsed token: %+v expiration error: %+v\n", jwtUser, err)
	}

}
