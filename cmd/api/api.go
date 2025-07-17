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

	mux.HandleFunc("POST /payments", payment.ProcessPaymentHandler)

	mux.HandleFunc("GET /payments-summary", payment.GetPaymentsSummaryHandler)

	config.Log.Info("starting server", "addr", config.Cfg.Addr)
	if err := http.ListenAndServe(config.Cfg.Addr, mux); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
