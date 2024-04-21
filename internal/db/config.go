package db

import "time"

type Config struct {
	Name              string        `koanf:"name"`
	IsSqlite          bool          `koanf:"is_sqlite"`
	IsPostgresql      bool          `koanf:"is_postgresql"`
	URL               string        `koanf:"url"`
	ConnectionTimeout time.Duration `koanf:"connection_timeout"`
}
