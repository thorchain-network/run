package runner

import (
	"net/http"
	"time"

	"run/client"
	"run/coincap"
	"run/config"
	"run/logger"

	"github.com/redis/go-redis/v9"
)

func Start(rdb *redis.Client, httpClient *http.Client, cfg config.App) {
	ticker := time.NewTicker(cfg.Frequency)
	defer ticker.Stop()

	for range ticker.C {
		if !client.RedisConnection(rdb) {
			logger.Log.Warn().Msg("Waiting for redis connection...")
			continue
		}

		if err := coincap.Fetch(rdb, httpClient, cfg); err != nil {
			logger.Log.Error().Err(err).Send()
		}
	}
}
