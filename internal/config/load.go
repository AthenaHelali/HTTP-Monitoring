package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	prefix    = "USERS"
	delimiter = "."
	separator = "__"
)

func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))
	return strings.ReplaceAll(base, separator, delimiter)
}

func New() Config {
	k := koanf.New(".")
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default:%s", err)
	}

	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Printf("error loading config.toml: %s", err)
	}
	if err := k.Load(env.Provider(prefix, delimiter, callbackEnv), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}
	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}
	fmt.Printf("%+v", instance)

	return instance

}
