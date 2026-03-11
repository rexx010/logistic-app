# ── Stage 1: Build ───────────────────────────────────────────────────
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o logisticApp ./main.go


# ── Stage 2: Run ─────────────────────────────────────────────────────
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/logisticApp .

EXPOSE 8080

CMD ["./logisticApp"]