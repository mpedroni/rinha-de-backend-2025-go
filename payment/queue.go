package payment

import (
	"sync"

	"github.com/mpedroni/rinha-backend-2025/config"
)

type Queue struct {
	// TODO: update to linked list
	payments []*Payment
	ch       chan *Payment
	m        sync.Mutex
	idx      int
}

func NewQueue() *Queue {
	q := &Queue{
		payments: make([]*Payment, 0, 100),
		ch:       make(chan *Payment),
	}

	go q.publish()
	return q
}

func (pq *Queue) publish() {
	for {
		pq.m.Lock()
		if pq.idx >= len(pq.payments) {
			pq.m.Unlock()
			continue
		}

		p := pq.payments[pq.idx]
		pq.idx++
		pq.ch <- p
		pq.m.Unlock()
	}
}

func (pq *Queue) Publish(payment *Payment) {
	pq.m.Lock()
	defer pq.m.Unlock()
	config.Log.Debug("publishing payment to queue", "payment", payment)
	pq.payments = append(pq.payments, payment)
}

func (pq *Queue) Subscribe() <-chan *Payment {
	return pq.ch
}
