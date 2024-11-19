# Stage 1: Build the Go application
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.work and go.work.sum files to set up the workspace
COPY go.work go.work.sum ./

# Copy the libs directory
COPY libs/pkg /app/libs/pkg

# Copy the user-management service
COPY services/user /app/services/user

# Set the working directory to the user-management service
WORKDIR /app/services/user

# Download dependencies
RUN go mod download

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-management main.go

# Stage 2: Create a minimal Docker image using Alpine
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/services/user/user-management .

# Ensure the binary is executable
RUN chmod +x ./user-management

# Expose the port your application will run on
EXPOSE 8080

# Run the executable
CMD ["./user-management"]

