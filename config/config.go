package config

import (
	"errors"
	"io"
	"log/slog"
	"os"
)

type Config struct {
	Addr        string
	Debug       bool
	ServiceName string
	Stdout      io.Writer
}

var Cfg *Config
var Log *slog.Logger

func Load() error {
	addr := os.Getenv("ADDR")
	debug := os.Getenv("DEBUG")
	service := os.Getenv("SERVICE_NAME")

	if addr == "" {
		addr = ":3000"
	}

	if service == "" {
		return errors.New("SERVICE_NAME environment variable is required")
	}

	isDebug := debug == "true"
	out := io.Discard
	if isDebug {
		out = os.Stderr
	}

	Cfg = &Config{
		Addr:        addr,
		Debug:       isDebug,
		ServiceName: service,
		Stdout:      out,
	}

	setupLogger()

	return nil
}
