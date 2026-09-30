package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Hostname     string `env:"APP_HOSTNAME" envDefault:"localhost"`
	OpenAIAPIKey string `env:"APP_OPENAI_API_KEY" envDefault:"***"`
	OpenAIModel  string `env:"APP_OPENAI_MODEL" envDefault:"gpt-5-mini"`
	MainBotToken string `env:"APP_MAIN_BOT_TOKEN" envDefault:"***"`
	DBDriver     string `env:"APP_DB_DRIVER" envDefault:"sqlite"`
	DBURL        string `env:"APP_DB_URL" envDefault:"sqlite.db"`
	MessagePrice int64  `env:"APP_MESSAGE_PRICE" envDefault:"50"`
}

func readConfig() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	log.Info().Msgf("APP_HOSTNAME: %s", config.Hostname)
	log.Info().Msgf("APP_OPENAI_API_KEY: (%d) bytes", len(config.OpenAIAPIKey))
	log.Info().Msgf("APP_OPENAI_MODEL: %s", config.OpenAIModel)
	log.Info().Msgf("APP_MAIN_BOT_TOKEN: (%d) bytes", len(config.MainBotToken))
	log.Info().Msgf("APP_DB_DRIVER: %s", config.DBDriver)
	log.Info().Msgf("APP_DB_URL: (%d) bytes", len(config.DBURL))
	log.Info().Msgf("APP_MESSAGE_PRICE: %d", config.MessagePrice)

	return &config, nil
}
