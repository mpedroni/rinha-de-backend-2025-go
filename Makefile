run:
	export SERVICE_NAME=api-1 && \
	export LOG_LEVEL=info && \
	export ADDR=:3002 && \
	export WORKERS_COUNT=4 && \
	export DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/rinha?sslmode=disable && \
	export PROCESSOR_DEFAULT_URL=http://localhost:8001 && \
	export PROCESSOR_FALLBACK_URL=http://localhost:8002 && \
	go build -o bin/api cmd/api/api.go && \
	./bin/api

up:
	docker compose -f ./infra/payment-processors/docker-compose.yaml up -d 
	docker compose up -d

down:
	docker compose -f ./infra/payment-processors/docker-compose.yaml down --remove-orphans
	docker compose down --remove-orphans

up-prod:
	docker compose -f ./infra/payment-processors/docker-compose.yaml up -d 
	docker compose -f docker-compose.prod.yaml up -d

down-prod:
	docker compose -f ./infra/payment-processors/docker-compose.yaml down --remove-orphans
	docker compose -f docker-compose.prod.yaml down --remove-orphans

db-connect:
	docker compose exec db psql -U postgres -d rinha

save-results:
	cp partial-results.json ./partial-results/$(shell date +%Y-%m-%d_%H-%M-%S).json