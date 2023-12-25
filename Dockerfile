# Start from the official Go image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin/api ./cmd/api

# Now, start a new stage with a smaller, minimal image
FROM alpine:latest

# Install the tzdata package
RUN apk --no-cache add ca-certificates tzdata

# Set the working directory and copy the binary from the previous stage
WORKDIR /root/
COPY --from=builder /app/bin/api .

# Command to run when the container starts
CMD ["./api"]
