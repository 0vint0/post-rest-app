    # Use a Go base image
    FROM golang:1.24-alpine

    # Add gcc for sql lite
    RUN apk add --no-cache gcc musl-dev linux-headers

    # Set the working directory inside the container
    WORKDIR /app

    # Copy go.mod and go.sum for dependency management
    COPY go.mod .
    COPY go.sum .

    # Download Go modules
    RUN go mod download

    # Copy the rest of the application source code
    COPY . .

    
    # Build the Go application
    # CGO_ENABLED keeping enabled for sqllite
    RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main .

    # Expose the port your GIN application listens on (e.g., 8080)
    EXPOSE 8080

    # Command to run the application
    CMD ["/app/main"]