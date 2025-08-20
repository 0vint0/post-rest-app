    # Use a Go base image
    FROM golang:1.24-alpine as build

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
    # add -v -x once to see progress during the long CGO compile
    RUN CGO_ENABLED=1 GOOS=linux go build -buildvcs=false -o /app/main .
    
        
    FROM alpine:3.20
     # no sqlite-libs needed in runtime when using bundled build
    RUN apk add --no-cache ca-certificates
    
    WORKDIR /app
    
    COPY --from=build /app/main /app/main
    
    EXPOSE 8080
    
    CMD ["/app/main"]
