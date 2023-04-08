# Use golang as the base image
FROM docker.io/golang:1.20.2-alpine AS builder

# Create a working directory
WORKDIR /app

# Copy all files from the current directory to the working directory
COPY . .

# Set goproxy
ENV GOPROXY=https://goproxy.cn,direct

# Build the application
RUN go build -o traded main.go

# Create the final Docker image
FROM docker.io/alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary file to the container
COPY --from=builder /app/traded /traded

# Set the entry command and arguments
ENTRYPOINT [ "/traded" ]