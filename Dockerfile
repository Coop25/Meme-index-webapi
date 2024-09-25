# Use the official Golang image as the base image
FROM golang:1.20-alpine AS builder

# Install Task
RUN apk add --no-cache curl \
    && curl -sL https://taskfile.dev/install.sh | sh
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

# Run Task commands
RUN task apigen
RUN task download-swagger-ui

RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /meme-index-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/meme-index-api"]