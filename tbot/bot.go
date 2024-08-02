package tbot

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/yanzay/tbot/v2"

	vmn "github.com/mrbelka12000/vps-monitoring-notifier"
	"github.com/mrbelka12000/vps-monitoring-notifier/config"
)

const (
	tickerNotifyInterval = 5 * time.Hour
)

type Bot struct {
	client *tbot.Client
	server *tbot.Server

	message chan vmn.ErrorMessage

	chatID string

	log *slog.Logger
}

func New(
	ctx context.Context,
	done chan<- bool,
	cfg config.Config,
	message chan vmn.ErrorMessage,
	opts ...opt,
) *Bot {

	telBot := tbot.New(cfg.TelegramToken)

	b := &Bot{
		client:  telBot.Client(),
		server:  telBot,
		message: message,
		chatID:  cfg.TelegramChatID,
		log:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	for _, opt := range opts {
		opt(b)
	}

	go b.sendMessages(ctx, done)

	return b
}

func (b *Bot) Start() error {
	return b.server.Start()
}

func (b *Bot) Stop() {
	b.server.Stop()
}

func (b *Bot) sendMessages(ctx context.Context, done chan<- bool) {
	defer func() {
		done <- true
	}()
	ticker := time.NewTicker(tickerNotifyInterval)

	for {
		select {
		case <-ticker.C:
			_, err := b.client.SendMessage(
				b.chatID,
				"All services working good, do not worry <3",
			)
			if err != nil {
				b.log.Error("send message to telegram bot failed", "error", err)
			}

		case msg := <-b.message:
			if len(msg) == 0 {
				b.log.Info("all services is okay")
				continue
			}

			_, err := b.client.SendMessage(
				b.chatID,
				generateServicesListMsg(msg),
			)
			if err != nil {
				b.log.Error("send message to telegram bot failed", "error", err)
			}
			ticker.Reset(tickerNotifyInterval)

		case <-ctx.Done():
			b.log.Info("sendMessages stopped successfully")
			return
		}
	}
}

func generateServicesListMsg(msg vmn.ErrorMessage) string {
	var result strings.Builder
	result.WriteString("Some services unavailable, please check\n\n")

	for i, v := range msg {
		result.WriteString(fmt.Sprintf("Service: %s has error: %s", v.Service, v.Msg))
		if i != len(msg)-1 {
			result.WriteString("\n\n")
		}
	}

	return result.String()
}
