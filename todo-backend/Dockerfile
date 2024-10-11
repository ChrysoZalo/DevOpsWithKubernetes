# Step 1: Build the Go binary in a temporary build container
FROM golang:1.22-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY *.go ./

# Build the Go application for Linux
RUN CGO_ENABLE=0 GOOS=linux go build -o /todo-backend
# Step 2: Create a minimal container to run the Go app
CMD ["/todo-backend"]
