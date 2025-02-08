FROM golang:1.23.5 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/api

FROM debian:bullseye-slim

COPY --from=builder /app/api /app/api

EXPOSE 8081

CMD ["/app/api"]
