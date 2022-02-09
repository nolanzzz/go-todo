# 1. Introduction
## 1.1 Project Introduction
> Go-TODO provides a backend api of a shared TODO List app based on Golang and the `Gin` framework. It also supports quick deployment with Docker.

Major api functions:
- User register/login
- Create/Update/Done/Undone TODO events
- All TODO events are public, however, only the creator can make changes. It's supposed to encourage users by seeing friends' progresses.
- Dynamically updated user rankings (by task numbers and total amount of minutes). Rankings are reset at midnight everyday.

Online api docs: [http://34.96.161.126/swagger/index.html](http://34.96.161.126/swagger/index.html)

Online api demo：GET - [http://34.96.161.126/api/v1/ranking/minutes/10](http://34.96.161.126/api/v1/ranking/minutes/10)

Testing `username`: `user1`

Testing password: `12345`

## 1.2 Tech Stacks
- Language：Golang
- Backend：RESTful api built with [Gin](https://gin-gonic.com)
- Database：
  - Using `MySQL`(8.0.21) as the main databaes
  - Using `Redis`(6.2.6) for recording the daily rankings
- ORM：Using Gorm v2(1.22.5) to implement basic data manipulation and also data migration
- Cache：Using `Redis` to stroe`JWT` tokens of active users，which supports further development on multipoint login, blacklists and so on
- API Docs：Using [Swagger](https://github.com/swaggo/swag) to auto-generate API documentation
- Config：Using [Viper](https://github.com/spf13/viper) to implement yaml configuration files
- Log：Using [zap](https://github.com/uber-go/zap) for logging

# 2. Getting Started

## 2.1 Setup

```bash
# Clone the repo
git clone https://github.com/nolanzzz/go-todo.git
cd go-todo
# Install dependencies
go generate
go build -o server main.go
# Run the binary executable and specify the config filename. Use config.yaml by default
./server -c config.yaml
```

## 2.2 Docker Deployment
This project includes a base `Dockerfile` as well as a `docker-compose` config file that includes `MySQL` and `Redis` services. Use Docker for quick deployment:
```bash
cd go-todo
docker-compose up
```

## 2.3 API Docs Auto-Generation using `Swagger`
The repo includes the most recent docs. Re-generate when need to make changes:
```bash
cd go-todo
swag init
```
Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in a browser to check generated docs after starting up the app.