package payment

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mpedroni/rinha-backend-2025/config"
)

func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req ProcessPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		config.Log.Error("failed to decode request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.Log.Debug("processing payment", "request", fmt.Sprintf("%+v", req))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`payment created`))
}

func GetPaymentsSummaryHandler(w http.ResponseWriter, r *http.Request) {
	var req GetPaymentsSummaryRequest
	req.From = r.URL.Query().Get("from")
	req.To = r.URL.Query().Get("to")

	config.Log.Debug("getting payments summary", "request", fmt.Sprintf("%+v", req))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`payments summary`))
}
