package config

import (
	"time"

	"github.com/spf13/viper"
)

var Configuration Config

type Config struct {
	Server struct {
		Mode            string        `mapstructure:"mode"`
		Port            int           `mapstructure:"port"`
		ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
		JWTSecret       string        `mapstructure:"jwt_secret"`
		TokenDuration   time.Duration `mapstructure:"token_duration"`
	} `mapstructure:"server"`

	Postgres struct {
		Host     string   `mapstructure:"host"`
		Port     int      `mapstructure:"port"`
		Database string   `mapstructure:"database"`
		Username string   `mapstructure:"username"`
		Password string   `mapstructure:"password"`
		Options  []string `mapstructure:"options"`
	} `mapstructure:"postgres"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	viper.SetConfigFile(path)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Configuration)
	return
}
