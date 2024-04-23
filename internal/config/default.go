package config

import "github.com/Milad75Rasouli/portfolio/internal/cipher"

func Default() Config {
	return Config{
		Debug: false,
		Cipher: cipher.Config{
			MoreSalt: "test123",
			Time:     1,
			Memory:   64 * 1024,
			Thread:   1,
			KeyLen:   64,
		},
	}
}
