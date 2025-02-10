# Use the official Golang image as the build stage
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to leverage Go's caching mechanism
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory to the container
COPY . .

# Build the API binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api/main.go

# Build the migrate binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o migrate ./cmd/migrate/main.go

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binaries from the builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .

# Copy the migrations and seeders
COPY db/migrations ./db/migrations
COPY db/seeders ./db/seeders

# Set the entrypoint for the container
ENTRYPOINT ["/app/server"]

# Optionally, expose the port your API server listens on
EXPOSE 8080
