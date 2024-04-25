package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshTokenExpireAfter = 48 // Hours
	AccessTokenExpireAfter  = 10 // Minute
)

var (
	InvalidJWTMethod   = errors.New("Invalid JWT Method")
	FailedToParseToken = errors.New("Failed to parse the token")
	InvalidToken       = errors.New("Invalid token")
	InvalidTokenClaims = errors.New("Invalid token claims")
	FailedToReadClaims = errors.New("Failed to read the claims")
)

type JWTUser struct {
	FullName     string
	Email        string
	Role         string
	InitiateTime time.Time
}
type JWTToken struct {
	RefreshToken RefreshJWT
	AccessToken  AccessJWT
}

func New(cfg Config) *JWTToken {
	return &JWTToken{
		RefreshToken: NewRefreshJWT(cfg),
		AccessToken:  NewAccessJWT(cfg),
	}
}

type RefreshJWT struct {
	SecretKey string
}

func NewRefreshJWT(c Config) RefreshJWT {
	return RefreshJWT{SecretKey: c.RefreshSecretKey}
}

// Caution: the jwtUser "InitiateTime" won't effect the final result. it's always time.Now()
func (j RefreshJWT) Create(jwtUser JWTUser) (string, error) {

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
func (j RefreshJWT) VerifyParse(tokenString string) (JWTUser, error) {
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
	initiateTime, err := claims.GetIssuedAt()
	if err != nil {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("expiration time claims error"))
	}
	jwtUser.InitiateTime = initiateTime.Time

	return jwtUser, nil
}

type AccessJWT struct {
	SecretKey string
}

func NewAccessJWT(cfg Config) AccessJWT {
	return AccessJWT{SecretKey: cfg.AccessSecretKey}
}

// Caution: the jwtUser "InitiateTime" won't effect the final result. it's always time.Now()
func (j AccessJWT) Create(jwtUser JWTUser) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     jwtUser.Email,
		"full_name": jwtUser.FullName,
		"role":      jwtUser.Role,
		"exp":       time.Now().Add(time.Minute * AccessTokenExpireAfter).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Caution: the jwtUser "InitiateTime" won't effect the final result. it's always time.Now()
func (j AccessJWT) VerifyParse(tokenString string) (JWTUser, error) {
	var jwtUser JWTUser

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwtUser, InvalidTokenClaims
	}

	var hasItem bool
	jwtUser.Email, hasItem = claims["email"].(string)
	if hasItem == false {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("email claims error"))
	}
	jwtUser.FullName, hasItem = claims["full_name"].(string)
	if hasItem == false {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("full name claims error"))
	}
	jwtUser.Role, hasItem = claims["role"].(string)
	if hasItem == false {
		return jwtUser, errors.Join(FailedToReadClaims, errors.New("role claims error"))
	}

	return jwtUser, nil
}
