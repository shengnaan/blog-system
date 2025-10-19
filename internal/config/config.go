package config

import (
	"os"
)

type Config struct {
	TZ string
}

func Load() Config {
	tz := os.Getenv("APP_TZ")
	if tz == "" {
		return Config{TZ: "UTC"}
	}
	return Config{TZ: tz}
}
