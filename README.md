# 数据库课设后端
一个用golang实现的学生管理系统的后端。使用gin-gonic框架、jwt验证、bcrypt加密密码

## 环境

- 开发环境：
    - macOS Catalina 10.15.4 19E287 x86_64
    - go version go1.13.7 darwin/amd64
    - MariaDB 5.5.64
    - Intel i7-8569U (8) @ 2.80GHz
    - 16 G Main Memory
- 运行环境： 为了在满足全校需求的情况下，系统能够流畅运行，建议使用
    - Ubuntu Server 18.04/CentOS 8-1905
    - go 1.13
    - MariaDB 5.5.64
    - AMD ryzen 9 3900x / Intel i9 10900k
    - 64G Main Memory

## 运行

1. 用 `extra`目录下的`.sql`文件创建数据库和数据库表。

2. 设置如下环境变量：
```shell
export DBURL="${数据库用户名}:${数据库密码}@/${数据库名}?charset=utf8&parseTime=True&loc=Local"
```
如果连接远程数据库，则设置：
```shell
export DBURL="${数据库用户名}:${数据库密码}@tcp(${服务器ip}:${端口号})/${数据库名}?charset=utf8&parseTime=True&loc=Local"
```

3. 在项目根目录下执行：
```shell
go build src/main.go
./main
```

