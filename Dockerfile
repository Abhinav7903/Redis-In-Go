# Use the official Golang image as the base
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules file (go.mod) and download dependencies
COPY go.mod ./

# Copy the application code into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

EXPOSE 1234

# Command to run the application
CMD ["./main"]
