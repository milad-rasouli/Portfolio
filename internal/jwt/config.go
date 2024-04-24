package jwt

type Config struct {
	SecretKey string `koanf:"secret_key"`
}
