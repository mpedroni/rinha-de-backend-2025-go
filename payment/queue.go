package payment

import (
	"sync"
)

type Queue struct {
	payments []ProcessPaymentRequest
	ch       chan ProcessPaymentRequest
	m        sync.Mutex
	idx      int
}

func NewQueue() *Queue {
	q := &Queue{
		payments: make([]ProcessPaymentRequest, 0, 100),
		ch:       make(chan ProcessPaymentRequest),
	}

	go q.publish()
	return q
}

func (pq *Queue) publish() {
	for {
		if pq.idx >= len(pq.payments) {
			continue
		}

		if len(pq.payments) > 90 {
			pq.m.Lock()

			pq.payments = pq.payments[pq.idx:]
			pq.idx = 0

			pq.m.Unlock()
		}

		req := pq.payments[pq.idx]
		pq.idx++
		pq.ch <- req
	}
}

func (pq *Queue) Publish(req ProcessPaymentRequest) {
	pq.m.Lock()
	defer pq.m.Unlock()
	pq.payments = append(pq.payments, req)
}

func (pq *Queue) Subscribe() <-chan ProcessPaymentRequest {
	return pq.ch
}
