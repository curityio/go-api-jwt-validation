Validate JWTs in Go
=============================================

.. image:: https://img.shields.io/badge/quality-experiment-red
    :target: https://curity.io/resources/code-examples/status/

.. image:: https://img.shields.io/badge/availability-source-blue
    :target: https://curity.io/resources/code-examples/status/

Set appropriate config in api/.env

Only works with RS256

```shell
docker build -t go-api . 
docker rm -f go-api
docker run --name go-api -p 8080:8080 go-api
```
