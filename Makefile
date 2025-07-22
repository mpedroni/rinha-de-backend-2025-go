run:
	SERVICE_NAME=api-local DEBUG=true ADDR=:3002 go run cmd/api/api.go

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