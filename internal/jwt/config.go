package jwt

type Config struct {
	RefreshSecretKey string `koanf:"refresh_secret_key"`
	AccessSecretKey  string `koanf:"access_secret_key"`
}
