FROM golang:alpine AS builder

LABEL maintainer="Muhamad Fitrah Ramadhan <mfitrahrmd>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

COPY .env .

RUN go get -d -v ./...

RUN go build -o app .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/app .

COPY --from=builder /app/.env .

EXPOSE 3000

CMD ["./app"]