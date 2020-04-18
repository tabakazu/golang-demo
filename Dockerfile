FROM golang:1.13-alpine

ENV GOPATH=/go
ENV GO111MODULE=on
RUN apk add --update --no-cache bash ca-certificates g++ gcc git mysql-client

ENV workdir /go/src/github.com/tabakazu
RUN mkdir -p ${workdir}
WORKDIR ${workdir}

# Usage
# $ docker build ./ -t golang-demo
# $ cd /your/path/pj-name
# $ docker run --rm -v $PWD:/go/src/github.com/tabakazu/pj-name -p 8090:8080 -it golang-demo /bin/bash
