FROM golang:1.17.6-alpine as BUILD

WORKDIR /tmp/bot

COPY . .

RUN apk update && \
    apk --no-cache --update add build-base git && \
    go mod download && \
    go mod verify && \
    go build cmd/bot/main.go

FROM alpine

COPY --from=BUILD /tmp/bot/build/. /home/bot/

ENTRYPOINT ./bot