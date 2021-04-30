FROM golang:1.16 as build

ENV GOPATH=/go
ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/tabakazu/webapi-app
WORKDIR /go/src/github.com/tabakazu/webapi-app
COPY . .

RUN apt-get update -qq && apt-get install -yq default-mysql-client
RUN go build -o /tmp/webapi-app/app
RUN GO111MODULE=off go get github.com/cosmtrek/air


FROM debian

COPY --from=build /tmp/webapi-app/app .

CMD ["./app"]
