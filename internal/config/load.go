package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

const (
	prefix    = "PORTFOLIO_"
	delimiter = "."
	separator = "__"
)

func New() Config {
	k := koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("Unable to load the default config: %s\n", err)
	}

	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Fatalf("Unable to load the config.toml: %s\n", err)
	}

	envCallback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))
		return strings.ReplaceAll(base, separator, delimiter)
	}
	if err := k.Load(env.Provider(prefix, delimiter, envCallback), nil); err != nil {
		log.Fatalf("Unable to load environment variables: %s\n", err)
	}

	var c Config
	if err := k.Unmarshal("", &c); err != nil {
		log.Fatalf("Unable to unmarshalling the config: %s\n", err)
	}
	return c
}
