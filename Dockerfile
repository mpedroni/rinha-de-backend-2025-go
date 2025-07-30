FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app -ldflags="-s -w" cmd/api/api.go

FROM alpine:latest AS prod

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /bin/app /app

USER nobody

ENTRYPOINT ["/app"]