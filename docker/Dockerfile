FROM golang:1.17.6-alpine AS build

ARG VERSION=dev

WORKDIR /tmp/bot

COPY . .

RUN go mod download && \
    go mod verify && \
    go build -ldflags="-X 'main.version=${VERSION}'" -o bot cmd/bot/main.go

FROM alpine

WORKDIR /home/bot

COPY --from=build /tmp/bot/bot /home/bot/bot

ENTRYPOINT ["./bot"]
