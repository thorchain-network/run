package coincap

import (
	"context"
	"encoding/json"
	"net/http"

	"run/config"
	"run/logger"

	"github.com/redis/go-redis/v9"
)

const URL = "https://api.coincap.io/v2/assets/"

type Data struct {
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}

type Response struct {
	Data Data `json:"data"`
}

func Fetch(rdb *redis.Client, httpClient *http.Client, cfg config.App) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", URL+cfg.Asset, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var coinData Response
	if err := json.NewDecoder(resp.Body).Decode(&coinData); err != nil {
		return err
	}

	dataJSON, err := json.Marshal(coinData.Data)
	if err != nil {
		return err
	}

	if err := rdb.Set(ctx, cfg.Redis.Prefix+"_"+cfg.Asset, dataJSON, 0).Err(); err != nil {
		return err
	}

	logger.Log.Info().Str(coinData.Data.Symbol, coinData.Data.PriceUsd).Msg(coinData.Data.Name)
	return nil
}
