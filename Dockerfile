# Use Ubuntu as the base image
FROM ubuntu:latest

# Set working directory
WORKDIR /app

# Install required dependencies
RUN apt update && apt install -y wget curl git unzip build-essential

# Install Go
RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz && \
    rm go1.21.6.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/app/go
ENV GOCACHE=/tmp/.cache/go-build
ENV GOMODCACHE=/tmp/.cache/go-mod

# Copy the entire project
COPY . .

# Install dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o go-books-api main.go

# Ensure the binary is executable
RUN chmod +x /app/go-books-api

# Expose the application port
EXPOSE 8080

# ✅ Run the Go API
CMD ["/app/go-books-api"]
