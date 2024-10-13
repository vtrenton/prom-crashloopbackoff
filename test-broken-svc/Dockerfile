# Use the specific Golang version with Alpine to build the application
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o /broken ./cmd/broken/broken.go

# Use the same image for the final stage to avoid additional layers
FROM golang:1.23.1-alpine

# Set the working directory inside the container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /kcgen .

# Execute and run the application at the entrypoint
ENTRYPOINT ["./broken"]

