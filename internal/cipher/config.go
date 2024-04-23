package cipher

type Config struct {
	MoreSalt string `koanf:"more_salt"`
	Time     uint32 `koanf:"time"`
	Memory   uint32 `koanf:"memory"`
	Thread   uint8  `koanf:"Thread"`
	KeyLen   uint32 `koanf:"KeyLen"`
}
