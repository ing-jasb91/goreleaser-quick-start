FROM golang:1.14-alpine AS builder

LABEL maintainer="Orlando Monteverde <orlmicron@gmail.com>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/app ./

FROM alpine:3.11.2
WORKDIR /root
COPY --from=builder /app/app .

ENV PORT=9000

EXPOSE ${PORT}
CMD [ "/root/app" ]
