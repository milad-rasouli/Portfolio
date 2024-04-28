package cipher

type Config struct {
	Paper  string `koanf:"paper"`
	Time   uint32 `koanf:"time"`
	Memory uint32 `koanf:"memory"`
	Thread uint8  `koanf:"Thread"`
	KeyLen uint32 `koanf:"KeyLen"`
}
