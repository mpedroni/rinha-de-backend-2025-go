package payment

import "time"

type PaymentStatus int

const (
	Pending PaymentStatus = iota
	Paid
	Failed
)

type PaymentProcessor int

const (
	Default PaymentProcessor = iota
	Fallback
)

type Payment struct {
	CorrelationID string
	Amount        int
	Status        PaymentStatus
	Processor     PaymentProcessor
	ReceivedAt    time.Time
	PaidAt        time.Time
}
