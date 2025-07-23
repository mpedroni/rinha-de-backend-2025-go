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
			for req := range wp.Queue.Subscribe() {
				config.Log.Debug("worker processing payment", "workerID", workerID, "request", req)

				err := wp.process(req)

				if err != nil {
					config.Log.Error("payment processing failed", "workerID", workerID, "correlationId", req.CorrelationID, "error", err)
					wp.Queue.Publish(req)
					continue
				}

				config.Log.Info("payment processed", "workerID", workerID, "correlationId", req.CorrelationID)
			}
		}(i)
	}
}

func (wp *WorkerPool) process(req ProcessPaymentRequest) error {
	config.Log.Debug("processing payment", "correlationId", req.CorrelationID, "amount", req.Amount)
	return nil
}
