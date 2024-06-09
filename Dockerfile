# base image for build
FROM golang:1.22-alpine as builder

# Set working dir
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Copy everything to working dir
COPY . .

# Build
RUN go build -o /server ./src/main.go

# Stage 2: Create the final image with the binary
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]