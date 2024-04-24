package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTokenStuff(t *testing.T) {
	cfg := Config{
		SecretKey: "VerySecretKey",
	}
	now := time.Now()
	user := model.User{
		ID:         1,
		FullName:   "foo bar",
		Email:      "bar@baz.com",
		Password:   "foobarbaz",
		IsGithub:   0,
		CreatedAt:  now,
		ModifiedAt: now,
		OnlineAt:   now,
	}

	jwt := NewJWT(cfg)
	token, err := jwt.createToken(user)
	assert.NoError(t, err)
	fmt.Printf("created token: %s", token)
}
