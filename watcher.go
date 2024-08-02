package vps_monitoring_notifier

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/mrbelka12000/vps-monitoring-notifier/config"
)

type (
	Watcher struct {
		services   map[serviceName]string
		servicesMu sync.Mutex

		log      *slog.Logger
		interval time.Duration

		messages chan ErrorMessage
	}

	pingFunc func(string) error

	errorResponse struct {
		Service serviceName
		Msg     string
	}

	ErrorMessage []errorResponse
)

func NewWatcher(
	ctx context.Context,
	done chan<- bool,
	cfg config.Config,
	opts ...opt) *Watcher {
	dur, err := time.ParseDuration(cfg.WatcherInterval)
	if err != nil {
		panic(err)
	}

	services := map[serviceName]string{
		transcripterBot: cfg.TranscripterBotURL,
		goalsScheduler:  cfg.GoalsSchedulerBotURL,
		mognoDB:         cfg.MongoURI,
		postgreSQL:      cfg.PGURI,
		mockServer:      cfg.MockServerURL,
		redisDB:         cfg.RedisAddr,
	}

	w := &Watcher{
		services: services,
		log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		interval: dur,
		messages: make(chan ErrorMessage, 10),
	}

	for _, opt := range opts {
		opt(w)
	}

	go w.run(ctx, done)

	return w
}

func (w *Watcher) GetMessages() chan ErrorMessage {
	return w.messages
}

func (w *Watcher) run(ctx context.Context, done chan<- bool) {
	defer func() {
		done <- true
	}()

	ticker := time.NewTicker(w.interval)
	for {
		select {

		case <-ticker.C:

			w.log.Info("executing service checker")
			w.servicesMu.Lock()
			required := map[serviceName]pingFunc{
				transcripterBot: ping,
				goalsScheduler:  ping,
				mognoDB:         pingMongo,
				postgreSQL:      pingPG,
				mockServer:      ping,
				redisDB:         pingRedis,
			}

			var errResp ErrorMessage

			for k, v := range required {
				if err := v(w.services[k]); err != nil {
					errResp = append(errResp, errorResponse{
						Service: k,
						Msg:     err.Error(),
					})
				}
			}

			w.servicesMu.Unlock()

			w.messages <- errResp

			w.log.Info("service checker done")

		case <-ctx.Done():
			close(w.messages)
			ticker.Stop()
			w.log.Info("watcher stopped successfully")
			return
		}
	}
}
