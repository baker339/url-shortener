# Use a minimal Go image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o server

# Expose the port
EXPOSE 3000

# Run the application
CMD ["/app/server"]