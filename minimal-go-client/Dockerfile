FROM golang:1.22-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
# including the Go modules files
# ignored files are listed in .dockerignore
COPY go.mod .
COPY go.sum .
COPY main.go .

# Download and install Go dependencies
RUN go mod download

# Build the Go application
RUN go build -a -o main main.go

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/main /app/
COPY content /app/content
COPY config /app/config

# Expose the port the application listens on
EXPOSE 8080

# Set the entry point of the container
CMD ["./main", "-pod", "-config", "config/config.yaml"]
