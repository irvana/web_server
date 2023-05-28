# Builder
FROM golang:1.19.7-alpine3.17 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /webserver

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o engine ./cmd/web/

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir -p /app/configs

WORKDIR /app

EXPOSE 8080

COPY --from=builder /webserver/engine /app
COPY ./configs/ /app/configs/

ENV SINARMAS_REDIS.ADDRESS="127.0.0.1:6379" \
    SINARMAS_SERVER.ADDRESS="127.0.0.1" \
    SINARMAS_SERVER.PORT=8080 \
    SINARMAS_AUTHENTICATION.PRIVPATH="configs/jwt.key" \
    SINARMAS_AUTHENTICATION.PUBPATH="configs/jwt.key.pub"

CMD /app/engine