FROM golang:1.17.6-alpine AS build

ARG VERSION=dev

WORKDIR /tmp/app

COPY ../../docker .

RUN go mod download && \
    go mod verify && \
    go build -ldflags="-X 'main.version=${VERSION}'" -o bot cmd/bot/bot.go

FROM alpine

WORKDIR /app

COPY --from=build /tmp/app/bot /app/bot

ENTRYPOINT ["./bot"]
