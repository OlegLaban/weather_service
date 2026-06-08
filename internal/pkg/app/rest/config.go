package rest

import (
	"errors"
	"fmt"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port int `yaml:"port"`
}

func NewFromYaml(b []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		return Config{}, errors.Join(ErrCantParseYamlDataForConfig, err)
	}

	return c, nil
}

func (c Config) Port() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}
