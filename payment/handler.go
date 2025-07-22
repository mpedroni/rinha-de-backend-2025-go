package payment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mpedroni/rinha-backend-2025/config"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req ProcessPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		config.Log.Error("failed to decode request", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ReceivedAt = time.Now()

	config.Log.Debug("processing payment", "request", fmt.Sprintf("%+v", req))

	if err := h.svc.SchedulePayment(r.Context(), req); err != nil {
		config.Log.Error("failed to process payment", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetPaymentsSummaryHandler(w http.ResponseWriter, r *http.Request) {
	var req GetPaymentsSummaryRequest
	req.From = r.URL.Query().Get("from")
	req.To = r.URL.Query().Get("to")

	config.Log.Debug("getting payments summary", "request", fmt.Sprintf("%+v", req))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`payments summary`))
}
