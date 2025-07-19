# Build
FROM golang:1.22-alpine AS builder

#Install git dan dep lain kalau perlu
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod & sum, dan download dependensi
COPY go.mod go.sum ./
RUN go mod download

#copy smeua file ke image
COPY . .

# build aplikasi
RUN go build -o main .

# Tahap 2: Buat image ringan hanya dengan binary
FROM alpine:latest

# Install SSL certs (untuk akses HTTPS, misal ke PostgreSQL)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary hasil build dari tahap pertama
COPY --from=builder /app/main .

# Port yang digunakan aplikasi (lihat di `EXPOSE 8080` pada Dockerfile sebelumnya)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]