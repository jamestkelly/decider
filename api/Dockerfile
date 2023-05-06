FROM golang:1.16.3-alpine3.13

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application inside the container
RUN go build -o app

# Expose the port that the application listens on
EXPOSE 8080

# Start the application when the container starts
CMD ["./app"]