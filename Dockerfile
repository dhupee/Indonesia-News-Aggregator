# Use the official Golang image as the base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the code into the container
COPY . .

# Install the required packages for loading environment variables from a file
RUN apk add --no-cache bash

# Build the Go application
RUN go build -o main .

# Run the application
CMD ["bash", "-c", "source .env && ./main"]