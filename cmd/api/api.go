package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mpedroni/rinha-backend-2025/config"
	"github.com/mpedroni/rinha-backend-2025/payment"
)

func main() {
	if err := config.Load(); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	dbconfig, err := pgxpool.ParseConfig(config.Cfg.DBConnectionString)
	if err != nil {
		panic(fmt.Errorf("failed to parse database connection string: %w", err))
	}

	dbconfig.MaxConns = 10
	dbconfig.MinIdleConns = 5

	db, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	mux := http.NewServeMux()

	queue := payment.NewQueue()
	svc := payment.NewService(queue, db, payment.Config{
		DefaultProcessorURL:  config.Cfg.DefaultProcessorURL,
		FallbackProcessorURL: config.Cfg.FallbackProcessorURL,
	})
	h := payment.NewHandler(svc)

	wp := &payment.WorkerPool{
		Num:     config.Cfg.WorkersCount,
		Queue:   queue,
		Service: svc,
	}
	wp.Run()

	mux.HandleFunc("POST /payments", h.ProcessPaymentHandler)
	mux.HandleFunc("GET /payments-summary", h.GetPaymentsSummaryHandler)

	config.Log.Info("starting server", "addr", config.Cfg.Addr)
	if err := http.ListenAndServe(config.Cfg.Addr, mux); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
