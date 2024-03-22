# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
//RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-pin

RUN go build -o /docker-gs-pin


# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]



# # Use the official Golang image as a base image
# FROM golang:1.19

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go modules configuration file
# COPY go.mod .
# COPY go.sum .

# # Download and install Go dependencies
# RUN go mod download

# # Copy the source code into the container
# COPY . .

# # Build the Go application
# RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

# # Use a lightweight base image to run the application
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the compiled binary from the builder stage
# COPY --from=builder /app/myapp .

# # Expose the port on which the application will listen
# EXPOSE 8083

# # Command to run the application
# CMD ["/myapp"]
