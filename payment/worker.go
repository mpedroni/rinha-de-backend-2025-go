package payment

import "github.com/mpedroni/rinha-backend-2025/config"

type WorkerPool struct {
	Num     int
	Queue   *Queue
	Service *Service
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.Num; i++ {
		go func(workerID int) {
			for payment := range wp.Queue.Subscribe() {
				config.Log.Debug("worker processing payment", "workerID", workerID, "request", payment)

				err := wp.process(payment)

				if err != nil {
					config.Log.Error("payment processing failed", "workerID", workerID, "correlationId", payment.CorrelationID, "error", err)
					wp.Queue.Publish(payment)
					continue
				}

				config.Log.Info("payment processed", "workerID", workerID, "correlationId", payment.CorrelationID)
			}
		}(i)
	}
}

func (wp *WorkerPool) process(p *Payment) error {
	config.Log.Debug("processing payment", "correlationId", p.CorrelationID, "amount", p.Amount)
	return nil
}
