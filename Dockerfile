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
    mkdir -p /app

WORKDIR /app

EXPOSE 9090

COPY --from=builder /webserver/engine /app

CMD /app/engine