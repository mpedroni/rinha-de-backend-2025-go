run:
	export SERVICE_NAME=api-1 && \
	export DEBUG=true && \
	export ADDR=:3002 && \
	export WORKERS_COUNT=4 && \
	export DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/rinha?sslmode=disable && \
	export PROCESSOR_DEFAULT_URL=http://localhost:8001 && \
	export PROCESSOR_FALLBACK_URL=http://localhost:8002 && \
	go run cmd/api/api.go

up:
	docker compose -f ./infra/payment-processors/docker-compose.yaml up -d 
	docker compose up -d

down:
	docker compose -f ./infra/payment-processors/docker-compose.yaml down --remove-orphans
	docker compose down --remove-orphans

db-connect:
	docker compose exec db psql -U postgres -d rinha

image-build:
	docker build -t mpedroni/rinha-backend-2025:latest .