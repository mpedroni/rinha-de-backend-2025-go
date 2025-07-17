package payment

type ProcessPaymentRequest struct {
	Amount        float64 `json:"amount"`
	CorrelationID string  `json:"correlationId"`
}

type GetPaymentsSummaryRequest struct {
	From string
	To   string
}
