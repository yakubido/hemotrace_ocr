package config

import (
	"strings"

	"github.com/spf13/viper"
)

const envPrefix = "OCR"

type Config struct {
	Debug bool
}

type ConsumerConfig struct {
}

type PublisherConfig struct {
}

func Load() (*Config, error) {
	var cfg Config

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	viper.SetConfigFile("config")
	viper.AddConfigPath("./")
	_ = viper.MergeInConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
