FROM golang:1.14.4-alpine

ENV GOPATH=/go
ENV GO111MODULE=on
RUN apk add --update --no-cache bash ca-certificates g++ gcc git mysql-client
RUN go get bitbucket.org/liamstask/goose/cmd/goose
