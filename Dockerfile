# ==========================
# Build Stage
# ==========================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Dibutuhkan jika ada dependency yang diambil dari git
RUN apk add --no-cache git

# Cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-s -w" \
    -o server \
    ./cmd/api

# ==========================
# Runtime Stage
# ==========================
FROM alpine:3.22

RUN apk add --no-cache \
    ca-certificates \
    tzdata

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
