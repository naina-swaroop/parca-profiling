# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies (none needed for standard library only)
RUN go mod download

# Copy source code
COPY . .

# Build the application with full debug symbols for profiling
# Explicitly preserve all debugging information
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags="all=-N -l" -ldflags="-extldflags=-static" -buildvcs=false -o main .

# Final stage - use distroless for better compatibility with profiling tools
FROM gcr.io/distroless/static-debian11

WORKDIR /

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 6060

# Command to run
CMD ["/main"]
