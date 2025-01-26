# Use the official Golang image as the base
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the application code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Create a new image to reduce the size
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder image
COPY --from=builder /app/main /app/

# Expose both HTTP and Telnet ports
EXPOSE 1234 5678

# Command to run the application
CMD ["./main"]
