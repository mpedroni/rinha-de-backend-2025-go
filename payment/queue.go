package payment

import (
	"sync"

	"github.com/mpedroni/rinha-backend-2025/config"
)

type Queue struct {
	// TODO: update to linked list
	payments []*Payment
	ch       chan *Payment
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{
		payments: make([]*Payment, 0, 100),
		ch:       make(chan *Payment),
	}
	q.cond = sync.NewCond(&q.mu)

	return q
}

func (q *Queue) Enqueue(payment *Payment) {
	config.Log.Debug("enqueueing payment", "payment", payment)
	q.mu.Lock()
	q.payments = append(q.payments, payment)
	q.cond.Signal()
	q.mu.Unlock()
}

func (q *Queue) Dequeue() *Payment {
	q.mu.Lock()
	for len(q.payments) == 0 {
		q.cond.Wait()
	}

	p := q.payments[0]
	q.payments = q.payments[1:]
	q.mu.Unlock()
	config.Log.Debug("dequeueing payment", "payment", p)
	return p
}

func (pq *Queue) Purge() {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	config.Log.Debug("purging payment queue")
	pq.payments = make([]*Payment, 0, 100)
}
