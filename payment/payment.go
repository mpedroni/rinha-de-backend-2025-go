package payment

import (
	"math"
	"time"
)

type PaymentProcessor int

const (
	Default PaymentProcessor = iota
	Fallback
)

type Payment struct {
	CorrelationID string
	Amount        int
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
	p.PaidAt = time.Now().UTC()
}
