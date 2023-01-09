package helper

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func YamlConverterHelper() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("config.yaml")
	if err != nil {
		panic(err)
	}
}
