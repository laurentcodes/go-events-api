# Start from the official Golang image for building
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o events-api main.go

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app

# Copy the built binary and any required files
COPY --from=builder /app/events-api ./
COPY api.db ./

# Expose the port (change if your app uses a different port)
EXPOSE 8080

# Run the binary
CMD ["./events-api"]