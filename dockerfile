# Stage 1: Builder
# Use the local platform for the builder to make builds faster (no emulation)
FROM --platform=$BUILDPLATFORM golang:1.24 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Argument for the target architecture (injected by Docker BuildKit)
ARG TARGETARCH

# Build the Go app for the target architecture
# CGO_ENABLED=0 for static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o assistant ./cmd/main.go

# Stage 2: Runner
# This will pull the alpine image for the TARGET platform automatically
FROM alpine:latest

WORKDIR /root/

# Install dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/assistant .

# Expose port
EXPOSE 8000

# Command to run
CMD ["./assistant"]