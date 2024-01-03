# 从零搭建一个 go web 服务 #

## 技术栈 ##
- HTTP框架：[gin](https://www.liwenzhou.com/posts/Go/gin/)
- API：[gin-swagger](https://www.liwenzhou.com/posts/Go/gin-swagger/)
- 数据库：
  - [go-redis](https://redis.uptrace.dev/zh/guide/go-redis.html)
  - [Postgresql](https://pkg.go.dev/gorm.io/driver/postgres@v1.5.2#section-documentation)
  - [Gorm](https://gorm.io/zh_CN/docs/index.html)
- 鉴权：[JWT](https://www.liwenzhou.com/posts/Go/json-web-token/)
- 读取配置：[Viper](https://www.liwenzhou.com/posts/Go/viper/)
- excel：[excelizelib](https://www.uiucode.com/view/52.html)
- 部署：Makefile
- 进程管理：[supervisor](http://supervisord.org/)


## 搭建应用 ##
因为我所使用的go version 大于1.11，因此这里使用go mod 进行项目管理

### 项目初始化 ###
```
mkdir -p /proj/stu-server
cd /proj/stu-server

# 生成 go.mod
go mod init stu-server
```

### 划分项目目录 ###
```
├── README.md
├── cmd
├── config                  # 配置相关
│   ├── config.go
│   ├── data_base.yaml
│   ├── model.go            # 配置相关的model
│   ├── stu.sql
│   └── stu_apiserver.yaml
├── docs                    # 主要存放swagger文档
├── go.mod
├── go.sum
├── internal                # 仅内部使用的包
│   ├── api                 # api 相关
│   │   ├── admin
│   │   ├── proxy
│   │   ├── student
│   │   └── tag
│   ├── dbsvc               # 数据库CRUD
│   ├── model               # 数据库｜Response 等结构体
│   ├── router              # 路由划分  
│   │   ├── router.go
│   │   └── router_api.go
│   └── websvc              # 业务代码
├── main.go                 # 入口文件
├── pkg                     # 可外部使用的包
│   ├── errors              # 错误日志处理
│   ├── excelizelib         # excel
│   ├── middleware          # 中间件
│   │   └── jwt.go          # jwt
│   └── util
│       ├── jwt.go
│       └── password.go
└── vendor                  # module
```


### 数据库 ###
#### PostgreSQL ####
##### 导出SQL #####
```
pg_dump -f xxx/stu.sql db_name
```
##### 导入SQL #####
```
psql -u username -h host -d dbname -f xxx.sql 
```

#### Redis ####


### 打包 ###
#### 简易打包 ####
```
# 默认以main.go 为入口文件，以go.mod module name 编译成二进制文件
go build 

# 运行打包的二进制，并关闭swagger输出，一般用于线上环境
./student-server -swagger=false
```

#### 使用makefile打包 ####
- [官方文档](https://makefiletutorial.com)
- [其他文档](https://eddycjy.gitbook.io/golang/di-3-ke-gin/makefile)

### supervisor ###
在项目中维护supervisor相关配置
supervisor/etc/supervisor/conf.d/xxx.conf
```
# 安装
apt install supervisor

# 启动服务
supervisorctl start service

# 重启服务
supervisorctl restart service

# 查看服务状态
supervisorctl status service

# supervisor 日志
less /var/log/supervisor/xxx.log

```