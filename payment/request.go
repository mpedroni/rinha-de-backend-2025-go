package payment

import "time"

type ProcessPaymentRequest struct {
	Amount        float64 `json:"amount"`
	CorrelationID string  `json:"correlationId"`
	ReceivedAt    time.Time
}

type GetPaymentsSummaryRequest struct {
	From string
	To   string
}
