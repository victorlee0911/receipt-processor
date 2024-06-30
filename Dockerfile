FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /receipt-processor

FROM builder AS tester
RUN go test -v ./...

FROM debian:buster-slim AS deployer

WORKDIR /

COPY --from=builder /receipt-processor .

EXPOSE 8080
CMD ["./receipt-processor"]
