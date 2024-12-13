# Use the official Go image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
