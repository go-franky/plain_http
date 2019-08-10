# Plain HTTP

A quick configuration for getting a server up an running with Go and Docker

## Prerequist

* docker
* docker-compose (optional)

## Usage

### Plain docker

```bash
docker build -t plain_http .
docker run -p 8080:8080 plain_http
```

### Using docker-compose

```bash
docker-compose build
docker-compose up
```

## AWS Lambda
You can use the webserver with API Gateway to configure a proxy to the lambda.
While this is not ideal since we don't benefit fully from API Gateway, it allows for small apps to be set up really quickly

To create the lambda file

```bash
make lambda
```

This will create a lambda.zip that is ready to be updated. And invoked.

To create the API Gateway, see the [API Gateway Configuration](./docs/aws-api-gateway.md)


