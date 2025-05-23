
# Use official Golang image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o operary main.go

# Environment
ENV PORT=8080
EXPOSE 8080

# Run
CMD ["./operary"]
