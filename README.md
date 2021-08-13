# go-api-jwt-validation

Set appropriate config in api/.env

Only works with RS256

```shell
docker build -t go-api . 
docker rm -f go-api
docker run --name go-api -p 8080:8080 go-api
```
