# Use golang as the base image
FROM golang:1.20.2-alpine AS builder

# Create a working directory
WORKDIR /app

# Copy all files from the current directory to the working directory
COPY . .

# Build the application
RUN go build -o trade main.go

# Create the final Docker image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary file to the container
COPY --from=builder /app/trade /app/trade

# Set the entry command and arguments
ENTRYPOINT [ "/app/trade" ]