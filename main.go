package main

import (
	"net/http"

	"run/client"
	"run/config"
	"run/runner"
)

func main() {
	cfg := config.Load()
	rdb := client.Redis(cfg.Redis.Url)
	httpClient := &http.Client{
		Timeout: cfg.Timeout,
	}

	go runner.Start(rdb, httpClient, cfg)

	select {}
}
