package config

import (
	"github.com/Milad75Rasouli/portfolio/internal/cipher"
	"github.com/Milad75Rasouli/portfolio/internal/db"
)

type Config struct {
	Debug    bool          `koanf:"debug"`
	Database db.Config     `koanf:"database"`
	Cipher   cipher.Config `koanf:"cipher"`
}
