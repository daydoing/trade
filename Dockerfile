# Use golang as the base image
FROM docker.io/golang:1.20.2-alpine AS builder

# Create a working directory
WORKDIR /app

# Copy all files from the current directory to the working directory
COPY . .

# Build the application
RUN go build -o trade main.go

# Create the final Docker image
FROM docker.io/alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary file to the container
COPY --from=builder /app/trade /trade
COPY --from=builder /app/traded.yaml /traded.yaml

# Set the entry command and arguments
ENTRYPOINT [ "/trade" ]