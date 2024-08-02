package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	vmn "github.com/mrbelka12000/vps-monitoring-notifier"
	"github.com/mrbelka12000/vps-monitoring-notifier/config"
	"github.com/mrbelka12000/vps-monitoring-notifier/tbot"
)

const (
	defaultTimeout = 5 * time.Second
)

func main() {

	mainCtx, mainCancel := context.WithCancel(context.Background())
	defer mainCancel()

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Get(mainCtx)
	if err != nil {
		log.Error(err.Error())
		return
	}

	http.DefaultClient.Timeout = defaultTimeout

	waitChan := make(chan bool, 1)
	w := vmn.NewWatcher(
		mainCtx,
		waitChan,
		cfg,
		vmn.WithLogger(log.With("instance", "watcher")),
	)

	bot := tbot.New(
		mainCtx,
		waitChan,
		cfg,
		w.GetMessages(),
		tbot.WithLogger(log.With("instance", "telegram_bot")),
	)

	gs := make(chan os.Signal)
	signal.Notify(gs, syscall.SIGINT, syscall.SIGTERM)
	stopChan := make(chan string)

	go func() {
		log.Info("Bot started")
		if err := bot.Start(); err != nil {
			stopChan <- err.Error()
			return
		}
	}()

	select {
	case sig := <-gs:
		log.Info("Received", "signal", sig)
		log.Info("Bot stopped properly")
		bot.Stop()
	case msg := <-stopChan:
		log.Info("Bot stopped", "error", msg)
	}
	mainCancel()

	<-waitChan
	<-waitChan

	close(waitChan)
	close(stopChan)
	close(gs)
}
