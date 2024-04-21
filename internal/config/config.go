package config

import "github.com/Milad75Rasouli/portfolio/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Database db.Config `koanf:"database"`
}
