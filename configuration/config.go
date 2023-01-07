package configuration

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	HttpServer HttpServerConfig `yaml:"http_server"`
}

type HttpServerConfig struct {
	Enabled         bool `yaml:"enabled" envconfig:"HTTP_SERVER_ENABLED"`
	Port            int  `yaml:"port" envconfig:"HTTP_SERVER_PORT"`
	RequestTimeout  int  `yaml:"request_timeout" envconfig:"HTTP_SERVER_REQUEST_TIMEOUT"`
	ResponseTimeout int  `yaml:"response_timeout" envconfig:"HTTP_SERVER_RESPONSE_TIMEOUT"`
}

func ReadYmlConfig(config *Config) (*Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		return config, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func ReadEnvConfig(config *Config) (*Config, error) {
	err := envconfig.Process("", config)
	return config, err
}
