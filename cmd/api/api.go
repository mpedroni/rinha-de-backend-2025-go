package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mpedroni/rinha-backend-2025/config"
	"github.com/mpedroni/rinha-backend-2025/payment"
)

func main() {
	if err := config.Load(); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	mux := http.NewServeMux()

	queue := payment.NewQueue()
	svc := payment.NewService(queue)
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
