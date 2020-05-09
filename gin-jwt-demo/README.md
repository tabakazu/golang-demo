# gin-jwt-demo

## Start web server

```bash
$ docker run --rm -v $PWD:/go/src/github.com/tabakazu/gin-jwt-demo -p 8090:8080 -it golang-demo /bin/bash
docker > $ cd gin-jwt-demo
docker > $ go run main.go
```

## Use API

```bash
# Login (Get token)
$ curl -X POST http://localhost:8090/login -H 'content-type: application/json' -d '{"username":"admin", "password":"admin"}'

# Call as an authorization user
$ curl -X GET http://localhost:8090/auth/hello -H 'content-type: application/json' -H 'Authorization: Bearer <トークン>'
```
