package util

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	QQ QQ `yaml:"qq"`
	MX MX `yaml:"mx"`
}

type QQ struct {
	Host  string `yaml:"host"`
	Group string `yaml:"group"`
}

type MX struct {
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(bytes, &config)
	return config, err
}
