package config

type Redis struct {
	Url    string
	Prefix string
}

func redisConfig(isProd bool) Redis {
	url := getEnv("REDIS_URL", "")
	if url == "" {
		url = "redis://localhost:6379/0"
		if isProd {
			url = "redis://redis:6379/0"
		}
	}

	return Redis{
		Url:    url,
		Prefix: getEnv("REDIS_PREFIX", "thormon_run"),
	}
}
