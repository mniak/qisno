package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Load() (config Config, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName(".pismo")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	config.OTP.Database = expandTilde(config.OTP.Database, home)
	return
}

func expandTilde(path, home string) string {
	if path == "~" {
		return home
	}
	if strings.HasPrefix(path, "~/") {
		return filepath.Join(home, path[2:])
	}
	return path
}
