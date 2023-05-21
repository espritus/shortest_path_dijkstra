# Start from a Go image.
FROM golang

# Set the working directory inside the container.
WORKDIR /app

# Copy the necessary Go modules files.
COPY go.mod go.sum ./

# Download and install Go module dependencies.
RUN go mod download

# Copy the entire project directory into the container.
COPY . .

# Build the Go application.
RUN go build -o main .

# Set the entry point for the container.
ENTRYPOINT ["./main"]
