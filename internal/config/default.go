package config

import (
	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/jwt"
)

func Default() Config {
	return Config{
		Debug:      false,
		AdminEmail: "changeIt@ChangeIt.com",
		Cipher: cipher.Config{
			Paper:  "changeIt",
			Time:   1,
			Memory: 64 * 1024,
			Thread: 1,
			KeyLen: 64,
		},
		JWT: jwt.Config{
			RefreshSecretKey: "changeIt",
			AccessSecretKey:  "changeIt",
		},
	}
}
