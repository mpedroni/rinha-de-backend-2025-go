package config

import (
	"log/slog"
)

func setupLogger() {
	opts := &slog.HandlerOptions{}

	if Cfg.Debug {
		opts.Level = slog.LevelDebug
	}

	Log = slog.
		New(slog.NewTextHandler(Cfg.Stdout, opts)).
		With("service", Cfg.ServiceName)
}
