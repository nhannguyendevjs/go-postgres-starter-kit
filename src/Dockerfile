# Use official Golang image
FROM golang:1.23.4

WORKDIR /app

# Copy everything
COPY . .

# Build the application
RUN go build -o main cmd/app/main.go

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
