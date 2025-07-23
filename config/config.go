package config

import (
	"errors"
	"io"
	"log/slog"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Addr                 string
	Debug                bool
	ServiceName          string
	WorkersCount         int
	DBConnectionString   string
	DefaultProcessorURL  *url.URL
	FallbackProcessorURL *url.URL

	Stdout io.Writer
}

var Cfg *Config
var Log *slog.Logger

func Load() error {
	addr := os.Getenv("ADDR")
	debug := os.Getenv("DEBUG")
	service := os.Getenv("SERVICE_NAME")
	workersCount := os.Getenv("WORKERS_COUNT")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	defaultProcessorURL := os.Getenv("DEFAULT_PROCESSOR_URL")
	fallbackProcessorURL := os.Getenv("FALLBACK_PROCESSOR_URL")

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

	if dbConnectionString == "" {
		return errors.New("DB_CONNECTION_STRING environment variable is required")
	}

	if defaultProcessorURL == "" {
		return errors.New("DEFAULT_PROCESSOR_URL environment variable is required")
	}

	if fallbackProcessorURL == "" {
		return errors.New("FALLBACK_PROCESSOR_URL environment variable is required")
	}

	isDebug := debug == "true"
	out := io.Discard
	if isDebug {
		out = os.Stderr
	}

	defaulUrl, err := url.Parse(defaultProcessorURL)
	if err != nil {
		return errors.New("invalid DEFAULT_PROCESSOR_URL")
	}
	fallbackUrl, err := url.Parse(fallbackProcessorURL)
	if err != nil {
		return errors.New("invalid FALLBACK_PROCESSOR_URL")
	}

	Cfg = &Config{
		Addr:                 addr,
		Debug:                isDebug,
		ServiceName:          service,
		WorkersCount:         workersCountInt,
		DBConnectionString:   dbConnectionString,
		DefaultProcessorURL:  defaulUrl,
		FallbackProcessorURL: fallbackUrl,

		Stdout: out,
	}

	setupLogger()

	return nil
}
