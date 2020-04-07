FROM golang:1.14-alpine
WORKDIR /app

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o app main.go

CMD ["./app"]
