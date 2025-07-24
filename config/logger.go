package config

import (
	"log/slog"
)

func setupLogger() {
	opts := &slog.HandlerOptions{}

	opts.Level = Cfg.LogLevel
	Log = slog.
		New(slog.NewTextHandler(Cfg.Stdout, opts)).
		With("service", Cfg.ServiceName)
}
