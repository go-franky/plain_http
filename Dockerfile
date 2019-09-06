FROM golang:latest AS builder
ARG GIT_REVISION
WORKDIR /app
COPY go.sum .
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/go-franky/plain_http/version.GitRevision=${GIT_REVISION}" -o server cmd/web/web.go 

FROM alpine
RUN apk --no-cache add ca-certificates curl
WORKDIR /app
COPY --from=builder ./app/server .
HEALTHCHECK --interval=5s --timeout=3s CMD curl --silent --fail http://localhost:8080/ || exit 1
EXPOSE 8080
ENTRYPOINT ["/app/server"]
