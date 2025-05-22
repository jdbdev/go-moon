# Build stage
FROM golang:1.21-alpine AS builder

# Add git for go mod download
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .
# Copy templates and static files
COPY --from=builder /app/web/templates ./web/templates

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"] 