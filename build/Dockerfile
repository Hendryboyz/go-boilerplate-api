# Stage 1: Build
FROM golang:1.23-alpine AS builder

ARG PRIVATE_GIT_REGISTRY

# Install required build tools
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download \
    && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o /app/bin/api

# Stage 2: Run
FROM alpine:3.18

# Install necessary packages for running the Go app
RUN apk add --no-cache ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/api /usr/local/bin/api

# Expose the application's port (adjust based on your app configuration)
EXPOSE 8081

# Command to run the web app
ENTRYPOINT ["api", "server"]
