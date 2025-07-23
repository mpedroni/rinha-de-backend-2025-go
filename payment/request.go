package payment

import "time"

type ProcessPaymentRequest struct {
	Amount        float64 `json:"amount"`
	CorrelationID string  `json:"correlationId"`
}

type GetPaymentsSummaryRequest struct {
	From string
	To   string
}

// payment processor

type PayRequestBody struct {
	CorrelationID string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	RequestedAt   string  `json:"requestedAt"`
}

func NewPayRequest(p *Payment) PayRequestBody {
	return PayRequestBody{
		CorrelationID: p.CorrelationID,
		Amount:        p.AmountAsFloat(),
		RequestedAt:   p.ReceivedAt.UTC().Format(time.RFC3339),
	}
}
