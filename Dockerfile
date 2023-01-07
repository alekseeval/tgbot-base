# syntax=docker/dockerfile:1

# BUILD stage
FROM golang:1.19.0-alpine3.15 as build-image

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build -o ./bin/bot ./cmd/main.go

# DEPLOY stage
FROM alpine:latest

WORKDIR /etc/tgbot

COPY --from=build-image /app/bin/bot /etc/tgbot/bot
COPY --from=build-image /app/config.yml /etc/tgbot/config.yml

CMD [ "./bot" ]
