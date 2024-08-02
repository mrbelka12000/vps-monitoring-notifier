package vps_monitoring_notifier

import (
	"log/slog"
)

type opt func(*Watcher)

func WithLogger(log *slog.Logger) opt {
	return func(w *Watcher) {
		w.log = log
	}
}
