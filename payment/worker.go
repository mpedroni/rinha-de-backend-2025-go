package payment

import (
	"context"

	"github.com/mpedroni/rinha-backend-2025/config"
)

type WorkerPool struct {
	Num     int
	Queue   *Queue
	Service *Service
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.Num; i++ {
		go func(workerID int) {
			for {
				// assuming it cannot be nil
				payment := wp.Queue.Dequeue()

				config.Log.Debug("worker processing payment", "workerID", workerID, "payment", payment)

				if err := wp.process(payment); err != nil {
					config.Log.Error("payment processing failed", "workerID", workerID, "correlationId", payment.CorrelationID, "error", err)
					wp.Queue.Enqueue(payment)
					continue
				}

				config.Log.Info("payment processed", "workerID", workerID, "correlationId", payment.CorrelationID)
			}
		}(i + 1)
	}
}

func (wp *WorkerPool) process(p *Payment) error {
	if err := wp.Service.Pay(context.TODO(), p); err != nil {
		return err
	}

	return nil
}
