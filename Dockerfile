# Use the official Go image
FROM golang:1.24.1

WORKDIR /app

COPY . .

# Install PostgreSQL development libraries
RUN apt-get update && apt-get install -y libpq-dev

# Set CGO flags so the compiler can find libpq-fe.h
ENV CGO_CFLAGS="-I/usr/include/postgresql"

RUN go mod tidy

RUN go build -o app

RUN mkdir -p /app/data

EXPOSE 8080
EXPOSE 6060
CMD ["/app/app"]
