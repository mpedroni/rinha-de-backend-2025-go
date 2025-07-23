package payment

import (
	"context"
	"time"
)

type Service struct {
	queue *Queue
}

func NewService(q *Queue) *Service {
	return &Service{
		queue: q,
	}
}

func (s *Service) SchedulePayment(ctx context.Context, req ProcessPaymentRequest) error {
	s.queue.Publish(&Payment{
		CorrelationID: req.CorrelationID,
		Amount:        int(req.Amount * 100),
		ReceivedAt:    time.Now(),
		Status:        Pending,
		Processor:     Default,
	})
	return nil
}
