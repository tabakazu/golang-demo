# HTTP server logging demo

## Start web server

```bash
$ docker run --rm -v $PWD:/go/src/github.com/tabakazu/http-server-logging-demo -p 8090:8080 -it golang-demo /bin/bash
docker > $ cd http-server-logging-demo
docker > $ go run main.go
```

## Use API

```bash
# Login (Get token)
$ curl -X POST http://localhost:8090/login -H 'content-type: application/json' -d '{"username":"admin", "password":"admin"}'

# Call as an authorization user
$ curl -X GET http://localhost:8090/auth/hello -H 'content-type: application/json' -H 'Authorization: Bearer <トークン>'
```
