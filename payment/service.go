package payment

import (
	"context"
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
	s.queue.Publish(req)
	return nil
}
