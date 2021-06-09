# validator demo

```bash
$ docker build ../ -t golang-demo
$ docker run --rm -v $PWD:/go/src/github.com/tabakazu/validator-demo -p 8090:8080 -it golang-demo /bin/bash

docker > cd validator-demo/
docker > go run main.go
```
