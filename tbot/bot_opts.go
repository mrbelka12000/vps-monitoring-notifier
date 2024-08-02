package tbot

import "log/slog"

type opt func(*Bot)

func WithLogger(log *slog.Logger) opt {
	return func(b *Bot) {
		b.log = log
	}
}
