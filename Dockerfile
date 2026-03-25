# Multi-stage build for optimal image size

# Stage 1: Builder
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum if they exist
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download 2>/dev/null || true

# Copy source code
COPY *.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o flinters-challenge

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS if needed
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/flinters-challenge ./

# Create result directory
RUN mkdir -p result

# Default to processing ad_data.csv if it exists in mounted volume
# Usage: docker run -v /path/to/data:/data myapp --input=/data/file.csv
ENTRYPOINT ["./flinters-challenge"]
CMD ["--input=/data/ad_data.csv"]
