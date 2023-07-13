# Start with a base Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o app ./cmd/main.go

# Expose the port that your application listens on
EXPOSE 8080

# Run the application
CMD ["./app"]
