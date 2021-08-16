Validate JWTs in Go
=============================================

[![Quality](https://img.shields.io/badge/quality-experiment-red)](https://curity.io/resources/code-examples/status/)
[![Availability](https://img.shields.io/badge/availability-source-blue)](https://curity.io/resources/code-examples/status/)

An example API in Go integrated with [JSON Web Token for Go](https://github.com/gbrlsnchs/jwt) to perform JWT verification and validation for API Authorization.

## Run the example

Set appropriate config in `api/.env`

Build and run the Docker image.

```shell
docker build -t go-api . 
docker run --name go-api -p 8080:8080 go-api
```

NOTE: The example uses the `RS256` algorithm. The code needs to be tweaked to support others.

## Documentation
This repository is documented and described in the [Securing a Go API with JWTs](https://curity.io/resources/learn/go-api/) article.

## Contributing

Pull requests are welcome. To do so, just fork this repo, and submit a pull request.

## License

The files and resources maintained in this repository are licensed under the Apache 2 license.

## More Information

Please visit curity.io for more information about the Curity Identity Server.

Copyright (C) 2021 Curity AB.
