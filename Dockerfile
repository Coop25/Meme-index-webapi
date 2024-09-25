# Use the official Golang image as the base image
FROM golang:1.20-alpine AS builder

# Install Task
RUN apk add --no-cache curl \
    && curl -sL https://taskfile.dev/install.sh | sh

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Run Task commands
RUN task apigen
RUN task download-swagger-ui

# Build the Go app
RUN go build -o main .

# Use a minimal base image to reduce the size of the final image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]