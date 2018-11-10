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
