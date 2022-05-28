FROM golang:1.18-alpine3.16 as build-env
RUN apk add --no-cache git gcc
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o openapi
FROM alpine:3.16
COPY --from=build-env /app/openapi .
EXPOSE 8080/tcp
USER 1001
ENTRYPOINT ["./openapi"]