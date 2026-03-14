# ----------- Stage 1: Build -----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy Go modules manifests for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o go-readme-stats .

# ----------- Stage 2: Final Image -----------
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from builder
COPY --from=builder /app/go-readme-stats .

# Set the entrypoint
ENTRYPOINT ["./go-readme-stats"]