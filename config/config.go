package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	PGURI    string `env:"PG_URI,required"`
	MongoURI string `env:"MONGO_URI,required"`

	TranscripterBotURL   string `env:"TRANSCRIPTER_BOT_URL,required"`
	GoalsSchedulerBotURL string `env:"GOALS_SCHEDULER_BOT_URL,required"`
	MockServerURL        string `env:"MOCK_SERVER_URL,required"`

	WatcherInterval string `env:"WATCHER_INTERVAL,default=30s"`

	TelegramToken string `env:"TELEGRAM_TOKEN,required"`

	TelegramChatID string `env:"TELEGRAM_CHAT_ID,required"`
}

func Get(ctx context.Context) (Config, error) {
	return parseConfig(ctx)
}

func parseConfig(ctx context.Context) (cfg Config, err error) {
	godotenv.Load()

	err = envconfig.Process(ctx, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("fill config: %w", err)
	}

	return cfg, nil
}
