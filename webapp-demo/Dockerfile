FROM golang:1.16 as build

ENV GOPATH=/go
ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/tabakazu/go-webapp
WORKDIR /go/src/github.com/tabakazu/go-webapp
COPY . .

RUN apt-get update -qq && apt-get install -yq default-mysql-client
RUN go build -o /tmp/go-webapp/app
RUN GO111MODULE=on go get \
  github.com/google/wire/cmd/wire
RUN GO111MODULE=off go get \
  github.com/cosmtrek/air \
  github.com/swaggo/swag/cmd/swag


FROM debian

COPY --from=build /tmp/go-webapp/app .

CMD ["./app"]
