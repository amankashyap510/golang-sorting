# Dockerfile
FROM golang:1.21.5

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 8000

CMD ["./app"]