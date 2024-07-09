# Start with a minimal Go image
FROM golang:alpine3.20 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY top_numbers.go .

# Build the Go application
RUN go build -o top_numbers .

# Start a new stage from scratch
FROM alpine:3.20

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary executable from the previous stage
COPY --from=builder /app/top_numbers .

# Set the entry point for the container
ENTRYPOINT ["./top_numbers"]
