FROM golang:1.17 AS builder

WORKDIR /todo
COPY . .

RUN go build -o app


EXPOSE 8080
#ENTRYPOINT ./app