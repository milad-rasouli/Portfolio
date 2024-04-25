package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var InvalidJWTMethod = errors.New("Invalid JWT Method")
var FailedToParseToken = errors.New("Failed to parse the token")
var InvalidToken = errors.New("Invalid token")
var InvalidTokenClaims = errors.New("Invalid token claims")
var FailedToReadClaims = errors.New("Failed to read the claims")

type JWT struct {
	cfg Config
}

func NewJWT(c Config) *JWT {
	return &JWT{cfg: c}
}

type JWTUser struct {
	FullName string
	Email    string
	Role     string
}

func (j JWT) CreateToken(jwtUser JWTUser) (string, error) {

	now := time.Now()
	claims := &jwt.RegisteredClaims{
		Issuer:    "MiladRasouli.ir",
		Subject:   jwtUser.FullName,
		Audience:  []string{jwtUser.Role},
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Second)), //TODO: choose a reasonable amount
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        jwtUser.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWT) VerifyParseUserToken(tokenString string) (JWTUser, error) {
	var jwtUser JWTUser

	token, err := jwt.ParseWithClaims(tokenString, new(jwt.RegisteredClaims),
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, InvalidJWTMethod
			}

			return []byte(j.cfg.SecretKey), nil
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
