package config

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	Addr         string
	Debug        bool
	ServiceName  string
	WorkersCount int

	Stdout io.Writer
}

var Cfg *Config
var Log *slog.Logger

func Load() error {
	addr := os.Getenv("ADDR")
	debug := os.Getenv("DEBUG")
	service := os.Getenv("SERVICE_NAME")
	workersCount := os.Getenv("WORKERS_COUNT")

	if addr == "" {
		addr = ":3000"
	}

	if service == "" {
		return errors.New("SERVICE_NAME environment variable is required")
	}

	if workersCount == "" {
		return errors.New("WORKERS_COUNT environment variable is required")
	}

	workersCountInt, err := strconv.Atoi(workersCount)
	if err != nil {
		return errors.New("WORKERS_COUNT must be an integer")
	}

	isDebug := debug == "true"
	out := io.Discard
	if isDebug {
		out = os.Stderr
	}

	Cfg = &Config{
		Addr:         addr,
		Debug:        isDebug,
		ServiceName:  service,
		WorkersCount: workersCountInt,
		Stdout:       out,
	}

	setupLogger()

	return nil
}
