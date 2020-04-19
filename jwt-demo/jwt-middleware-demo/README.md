# jwt-middleware-demo

## Start web server

```bash
$ docker run --rm -v $PWD:/go/src/github.com/tabakazu/jwt-middleware-demo -p 8090:8080 -it golang-demo /bin/bash
docker > $ cd jwt-middleware-demo
docker > $ go run main.go
```

## Use API

```bash
# Get token
$ curl -X GET http://localhost:8090/auth -H 'content-type: application/json'
# Get user info with token
$ curl -X GET http://localhost:8090/user -H 'content-type: application/json' -H 'Authorization: Bearer <トークン>'
{"name":"Taro Yamada"}
```
