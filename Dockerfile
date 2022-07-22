FROM golang:alpine

WORKDIR /app/test-go

RUN apk update && apk add git upx curl vim busybox-extras

RUN go mod init test-go