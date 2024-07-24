FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/jutionck/golang-todo-apps/
WORKDIR /github.com/jutionck/golang-todo-apps/

RUN go mod tidy
RUN go build -o ./.bin/app .

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/jutionck/golang-todo-apps/.bin/app .
COPY --from=0 /github.com/jutionck/golang-todo-apps/.env .

CMD ["./app"]