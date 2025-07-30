package payment

import (
	"math"
	"time"
)

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

func ParseMoney(amount float64) int {
	return int(math.Round(amount * 100))
}

func (p *Payment) AmountAsFloat() float64 {
	return float64(float64(p.Amount) / 100.0)
}

func (p *Payment) Paid() {
	p.Status = Paid
	p.PaidAt = time.Now().UTC()
}
