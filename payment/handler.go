package payment

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	config.Log.Debug("processing payment", "request", fmt.Sprintf("%+v", req))

	if err := h.svc.SchedulePayment(r.Context(), req); err != nil {
		config.Log.Error("failed to process payment", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.Log.Info("payment scheduled successfully", "request", fmt.Sprintf("%+v", req))

	w.WriteHeader(http.StatusAccepted)
}

func (h *Handler) GetPaymentsSummaryHandler(w http.ResponseWriter, r *http.Request) {
	var req GetPaymentsSummaryRequest
	req.From = r.URL.Query().Get("from")
	req.To = r.URL.Query().Get("to")

	config.Log.Debug("getting payments summary", "request", fmt.Sprintf("%+v", req))

	summary, err := h.svc.GetPaymentsSummary(r.Context(), req)
	if err != nil {
		config.Log.Error("failed to get payments summary", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(summary)
	if err != nil {
		config.Log.Error("failed to marshal payments summary", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(body); err != nil {
		config.Log.Error("failed to write payments summary", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.Log.Info("payments summary retrieved successfully", "req", req, "summary", fmt.Sprintf("%+v", summary))
}

func (h *Handler) PurgePaymentsHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.PurgePayments(r.Context()); err != nil {
		config.Log.Error("failed to purge payments", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
