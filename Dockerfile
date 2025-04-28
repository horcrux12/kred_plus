# Gunakan base image golang
FROM golang:1.23-alpine

# Set environment variable
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Buat folder kerja di container
WORKDIR /app

# Copy go.mod dan go.sum, lalu download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy semua file project
COPY . .

# Build aplikasi
RUN go build -o main .

# Port yang akan dibuka
EXPOSE 8080

# Command untuk menjalankan aplikasi
CMD ["./main", "prod"]
