package jwt

import (
	"time"

	"github.com/Milad75Rasouli/portfolio/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	cfg Config
}

func NewJWT(c Config) *JWT {
	return &JWT{cfg: c}
}

func (j JWT) createToken(user model.User) (string, error) {

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer:    "MiladRasouli.ir",
		Subject:   user.FullName,
		Audience:  []string{"admin"},                      //TODO: change based on their individual role
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)), //TODO: choose a reasonable amount
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
