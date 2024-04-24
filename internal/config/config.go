package config

import (
	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/db"
	"github.com/Milad75Rasouli/portfolio/internal/jwt"
)

type Config struct {
	Debug    bool          `koanf:"debug"`
	Database db.Config     `koanf:"database"`
	Cipher   cipher.Config `koanf:"cipher"`
	JWT      jwt.Config    `koanf:"jwt"`
}
