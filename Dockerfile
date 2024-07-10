# Gunakan image resmi Golang sebagai base image
FROM golang:1.21

# Set environment variables
ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Buat direktori kerja untuk aplikasi
WORKDIR /app

# Copy file go.mod dan go.sum jika ada
COPY go.mod ./
RUN go mod tidy

# Copy semua file ke direktori kerja
COPY . .

# Copy file .env
COPY .env .env

# Compile aplikasi
RUN go build -o main .

# Eksekusi aplikasi
CMD ["./main"]
