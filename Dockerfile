FROM golang:1.16

ENV GOPATH=/go
ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -yq default-mysql-client
ENV workdir /go/src/github.com/tabakazu
RUN mkdir -p ${workdir}
WORKDIR ${workdir}

# Usage
# $ docker build ./ -t golang-demo
# $ cd /your/path/pj-name
# $ docker run --rm -v $PWD:/go/src/github.com/tabakazu/pj-name -p 8090:8080 -it golang-demo /bin/bash
