# Stage 1: Builder
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# CGO_ENABLED=0 is required for alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /assistant ./cmd/main.go

# Stage 2: Runner
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates and tzdata for HTTPS and timezones
RUN apk --no-cache add ca-certificates tzdata

# Copy the Pre-built binary from the previous stage
COPY --from=builder /assistant .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./assistant"]