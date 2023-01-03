FROM golang:1.19.2-alpine3.16 AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/worlder
ENV GOSUMDB=off
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /go/bin/worlder main.go

FROM alpine

RUN apk add --no-cache tzdata ca-certificates libc6-compat
ENV TZ Asia/Jakarta

WORKDIR /go/bin/worlder

COPY --from=builder /go/bin/worlder /go/bin/worlder/worlder
COPY --from=builder /go/src/worlder/.env.example /go/src/worlder/.env

ENTRYPOINT ["/go/bin/worlder/worlder", "-migrate"]
