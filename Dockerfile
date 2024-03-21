# Use the official Golang image as a base image
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules configuration file
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Use a lightweight base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port on which the application will listen
EXPOSE 8083

# Command to run the application
CMD ["./myapp"]
