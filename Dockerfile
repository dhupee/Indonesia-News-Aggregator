# Base image
FROM golang:1.21.5-bookworm

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code
COPY . .

# Install the dependencies
RUN apt update && apt upgrade -y
RUN go mod download
RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps

# Build the application
RUN go build -o ./appbin ./main.go

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./appbin"]