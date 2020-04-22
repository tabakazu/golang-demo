# gorm-demo

## Use container

```bash
$ docker-compose run --rm app bash
docker > cd gorm-demo/

# Check connection
docker > $ go run main.go

# Migrate schema
docker > $ go run migrate.go

# Insert sample data
docker > $ go run insert_seed.go
```

## Login MySQL container

```bash
$ docker-compose run --rm app bash
docker > mysql -uroot -proot -h mysql
```
