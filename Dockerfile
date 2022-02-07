FROM golang:1.17 AS builder

WORKDIR /todo
COPY . .

RUN go build -o app && curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz


EXPOSE 8080
#ENTRYPOINT ./app