## 1. 介绍
### 1.1 项目简介
> Go-TODO提供一个基于Go语言和`Gin`框架开发的共享TODO待办事项管理软件的后端api，支持Docker快速部署

主要api功能：
- 用户注册/登录
- 新建/修改/完成TODO事项
- TODO事项所有人可见，只有创建者有权修改，达到与小伙伴们互相督促的效果
- 动态更新的当日完成事项数和总完成时间的排行，零点清零，同样达到与小伙伴们互相督促的效果

在线api文档: [http://34.96.161.126/swagger/index.html](http://34.96.161.126/swagger/index.html)

在线测试示例：GET - [http://34.96.161.126/api/v1/ranking/minutes/10](http://34.96.161.126/api/v1/ranking/minutes/10)

测试用户名：user1

测试密码：12345

### 1.2 技术选型
- 语言：Golang
- 后端：用 [Gin](https://gin-gonic.com) 快速搭建Restful风格的API
- 数据库：
  - 使用`MySQL`(8.0.21)为主数据库
  - 使用`Redis`(6.2.6)记录当日用户排行数据
- ORM：使用Gorm v2(1.22.5)实现对数据库的基本操作以及数据迁移
- 缓存：使用`Redis`实现记录当前登录用户的`JWT`令牌，可供开发多点登录限制以及黑名单等操作
- API文档：使用 [Swagger](https://github.com/swaggo/swag) 构建自动化文档
- 配置文件：使用 [Viper](https://github.com/spf13/viper) 实现yaml格式的配置文件
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录

## 2. 使用说明

### 2.1 项目运行

```bash
# 克隆项目
git clone https://github.com/nolanzzz/go-todo.git
cd go-todo
# 使用 go mod 并安装go依赖包
go generate
go build -o server main.go
# 运行二进制，并注明要使用的配置文件，不注明则默认使用config.yaml
./server -c config.yaml
```

### 2.2 Docker部署
本项目含Dockerfile和包含`MySQL`以及`Redis`的docker-compose配置文件，可使用Docker一键快速部署：
```bash
cd go-todo
docker-compose up
```

### 2.3 `Swagger`自动化API文档
项目已含最新docs，若需修改并重新生成API文档：
```bash
cd go-todo
swag init
```
启动项目后可前往 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) 查看生成的文档