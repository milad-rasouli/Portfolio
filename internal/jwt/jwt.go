package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshTokenExpireAfter = 48 // Hours
)

var (
	InvalidJWTMethod   = errors.New("Invalid JWT Method")
	FailedToParseToken = errors.New("Failed to parse the token")
	InvalidToken       = errors.New("Invalid token")
	InvalidTokenClaims = errors.New("Invalid token claims")
	FailedToReadClaims = errors.New("Failed to read the claims")
)

type RefreshJWT struct {
	SecretKey string
}

func NewRefreshJWT(c Config) *RefreshJWT {
	return &RefreshJWT{SecretKey: c.RefreshSecretKey}
}

type JWTUser struct {
	FullName string
	Email    string
	Role     string
}

func (j RefreshJWT) CreateRefreshToken(jwtUser JWTUser) (string, error) {

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer:    "MiladRasouli.ir", //TODO: take it from the config
		Subject:   jwtUser.FullName,
		Audience:  []string{jwtUser.Role},
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * RefreshTokenExpireAfter)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jwtUser.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j RefreshJWT) VerifyParseRefreshToken(tokenString string) (JWTUser, error) {
	var jwtUser JWTUser

	token, err := jwt.ParseWithClaims(tokenString, new(jwt.RegisteredClaims),
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, InvalidJWTMethod
			}

			return []byte(j.SecretKey), nil
		})

	if err != nil {
		return jwtUser, errors.Join(FailedToParseToken, err)
	}
	if token.Valid == false {
		return jwtUser, InvalidToken
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if ok == false {
		return jwtUser, InvalidTokenClaims
	}

	jwtUser.Email = claims.ID
	s, err := claims.GetSubject()
	if err != nil {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("subject claims error"))
	}
	jwtUser.FullName = s
	a, err := claims.GetAudience()
	if err != nil {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("audience claims error"))
	}
	jwtUser.Role = strings.Join(a, "")

	return jwtUser, nil
}
