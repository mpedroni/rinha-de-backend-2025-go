package payment

import (
	"sync"
)

type Queue struct {
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
		if pq.idx >= len(pq.payments) {
			continue
		}

		if len(pq.payments) > 90 {
			pq.m.Lock()

			pq.payments = pq.payments[pq.idx:]
			pq.idx = 0

			pq.m.Unlock()
		}

		p := pq.payments[pq.idx]
		pq.idx++
		pq.ch <- p
	}
}

func (pq *Queue) Publish(payment *Payment) {
	pq.m.Lock()
	defer pq.m.Unlock()
	pq.payments = append(pq.payments, payment)
}

func (pq *Queue) Subscribe() <-chan *Payment {
	return pq.ch
}
