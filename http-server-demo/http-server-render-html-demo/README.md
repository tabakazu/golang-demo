# Render html with HTTP server

## Start web server

```bash
$ docker run --rm -v $PWD:/go/src/github.com/tabakazu/http-server-render-html-demo -p 8090:8080 -it golang-demo /bin/bash
docker > $ cd http-server-render-html-demo
docker > $ go run main.go
```
