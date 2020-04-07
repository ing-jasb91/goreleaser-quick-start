FROM golang:1.14-alpine AS builder
WORKDIR /app

COPY . .
RUN go build -o app .

FROM alpine:3.11.2
WORKDIR /root
COPY --from=builder /app/app ./
CMD ["./app"]

