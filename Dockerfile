FROM golang:latest AS builder
WORKDIR /app
COPY go.sum .
COPY go.mod .
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/web/web.go 

FROM alpine
RUN apk --no-cache add ca-certificates curl
WORKDIR /app
COPY --from=builder ./app/server .
HEALTHCHECK --interval=5s --timeout=3s CMD curl --silent --fail http://localhost:8080/ || exit 1
EXPOSE 8080
ENTRYPOINT ["/app/server"]
