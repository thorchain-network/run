package config

import (
	"time"
)

type App struct {
	Environment string
	Asset       string
	Redis       Redis
	Frequency   time.Duration
	Timeout     time.Duration
}

func Load() App {
	environment := getEnv("APP_ENV", "development")
	isProd := environment == "production"
	frequency := time.Duration(getEnvAsInt("FREQUENCY", 4)) * time.Second

	return App{
		Environment: environment,
		Asset:       getEnv("ASSET", "thorchain"),
		Redis:       redisConfig(isProd),
		Frequency:   frequency,
		Timeout:     frequency * 9 / 10,
	}
}
