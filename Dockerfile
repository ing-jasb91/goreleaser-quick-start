FROM golang:1.14-alpine AS builder
WORKDIR /app
ENV GO111MODULE=on
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM alpine:3.11.2
WORKDIR /root
COPY --from=builder /app/app ./
CMD ["./app"]

