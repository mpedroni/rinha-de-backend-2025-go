package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	queue  *Queue
	client http.Client
	cfg    Config
	db     *pgxpool.Pool
}

type Config struct {
	DefaultProcessorURL  *url.URL
	FallbackProcessorURL *url.URL
}

var (
	ErrPaymentAlreadyProcessed = errors.New("payment already processed")
)

func NewService(q *Queue, db *pgxpool.Pool, cfg Config) *Service {
	return &Service{
		queue: q,
		// TODO: set timeout?
		client: http.Client{},
		cfg:    cfg,
		db:     db,
	}
}

func (s *Service) SchedulePayment(ctx context.Context, req ProcessPaymentRequest) error {
	s.queue.Publish(&Payment{
		CorrelationID: req.CorrelationID,
		Amount:        ParseMoney(req.Amount),
		ReceivedAt:    time.Now(),
		Status:        Pending,
		Processor:     Default,
	})
	return nil
}

func (s *Service) Pay(ctx context.Context, p *Payment) error {
	if err := s.pay(ctx, p); err != nil {
		return err
	}

	p.Paid()

	if err := s.persist(ctx, p); err != nil {
		return err
	}

	return nil
}

func (s *Service) pay(ctx context.Context, p *Payment) error {
	url := s.cfg.DefaultProcessorURL
	if p.Processor == Fallback {
		url = s.cfg.FallbackProcessorURL
	}

	body, err := json.Marshal(NewPayRequest(p))
	if err != nil {
		return fmt.Errorf("failed to marshal payment request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String()+"/payments", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create payment request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnprocessableEntity {
		return ErrPaymentAlreadyProcessed
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("payment processor returned status %d", resp.StatusCode)
	}

	return nil
}

func (s *Service) persist(ctx context.Context, p *Payment) error {
	_, err := s.db.Exec(ctx, "INSERT INTO payments (correlation_id, amount, received_at, status, processor, paid_at) VALUES ($1, $2, $3, $4, $5, $6)",
		p.CorrelationID, p.Amount, p.ReceivedAt.UTC(), p.Status, p.Processor, p.PaidAt.UTC())
	if err != nil {
		return fmt.Errorf("failed to persist payment: %w", err)
	}

	return nil
}
