# Gunakan image Go resmi sebagai base image
FROM golang:1.24-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum (jika ada) ke working directory
COPY go.mod go.sum ./

# Download dependensi Go
RUN go mod download

# Copy seluruh kode aplikasi ke working directory
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Gunakan image Alpine kecil untuk production
FROM alpine:latest

# Install dependensi yang diperlukan (jika ada)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy binary yang sudah di-build dari stage builder
COPY --from=builder /app/main .

# Expose port yang digunakan oleh aplikasi
EXPOSE 8081

# Jalankan aplikasi
CMD ["./main"]