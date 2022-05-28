FROM golang:1.18-alpine AS build

ARG VERSION=dev

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-X 'main.version=${VERSION}'-w -s" -o bot cmd/main.go

FROM alpine

COPY --from=build /build/bot /opt/bot/bot

ENTRYPOINT ["/opt/bot/bot"]
