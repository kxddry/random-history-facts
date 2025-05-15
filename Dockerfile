FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal ./internal/
COPY config ./config/
COPY migrations ./migrations/

RUN go build -o migrator ./cmd/migrator
RUN go build -o server ./cmd/random-history-facts

EXPOSE 8080

CMD ["sh", "-c", "./migrator --config=./config/migrations.yaml && ./server"]
