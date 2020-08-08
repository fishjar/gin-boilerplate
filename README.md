# Gin boilerplate

基于 `Gin` 封装的样板项目

## 项目特性

- 基础框架 `gin`： github.com/gin-gonic/gin
  - 跨域中间件：github.com/gin-contrib/cors
- 数据库 orm 框架 `gorm`： github.com/jinzhu/gorm
  - 自动创建表，插入初始数据
  - 支持事务及原始 SQL 查询
  - 支持查询缓存（待优化）：github.com/8treenet/gcache
- 多环境配置（yml 文件）依赖包：github.com/jinzhu/configor
- `jwt` 登录验证功能，依赖包：github.com/dgrijalva/jwt-go
  - 认证数据缓存在 redis
  - 可依据路由配置角色权限校验
- 模型定义中 `validator.v8` 校验： godoc.org/gopkg.in/go-playground/validator.v8
- `swagger` 文档生成及访问：github.com/go-openapi/swag
  - 访问支持 BasicAuth
- 日志依赖包：github.com/sirupsen/logrus
  - 日志分割：github.com/lestrrat-go/file-rotatelogs
- `redis` 依赖包：github.com/go-redis/redis/v7
- 单文件/多文件上传功能
  - 文件名 hash 处理
  - 如果是图片，自动生成缩略图，依赖包：github.com/disintegration/imaging
- 分布式任务队列：github.com/hibiken/asynq
  - 基于 redis
- 定时任务：github.com/robfig/cron
- 分布式业务锁：github.com/bsm/redislock
  - 基于 redis
- 基于 docker 部署
- 每个模型实现 8 个基本接口
  - 分页查询（GET）
  - 单条查询（GET）
  - 创建单条（POST）
  - 更新单条（PATCH）
  - 删除单条（DELETE）
  - 查询或创建单条（POST）
  - 批量更新（PATCH）
  - 批量删除（DELETE）

## 开发指引

```sh
# 确保已安装go，及$GOPATH环境变量已配置
# Go 1.13 and above
go version
echo $GOPATH

# 设置代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 创建并进入目录
mkdir -p $GOPATH/src/github.com/fishjar/gin-boilerplate && cd "$_"

# 克隆项目
git clone https://github.com/fishjar/gin-boilerplate.git .

# 确认配置文件(尤其数据库相关配置)
vi config/config.go
vi config/config.yml
vi config/config.development.yml

# 如有需要，运行下列命令启动一个mysql数据库服务
docker-compose -f docker-compose-mysql.yml up -d

# 如有需要，启动redis
docker-compose -f docker-compose-redis.yml up -d

# 安装依赖
go get

# 生成Swagger文档
swag init

# 开发启动
go run main.go

# 访问swagger文档
http://localhost:4000/swagger/index.html

# 测试：登录
curl -X POST http://localhost:4000/admin/account/login \
-H "Content-Type: application/json" \
-d '{"username":"gabe","password":"123456"}' | python -m json.tool

# 测试：查询记录，注意替换<token>为实际值
curl http://localhost:4000/admin/users \
-H "Authorization: Bearer <token>" | python -m json.tool
```

## 部署指引

### 方案一

完整的 `golang` 镜像中编译运行

```sh
docker-compose -f docker-compose.yml up -d
```

### 方案二

`golang:alpine` 镜像中编译运行

```sh
docker-compose -f docker-compose-alpine.yml up -d
```

### 方案三（镜像最小）

本地编译，`alpine:latest` 镜像中运行

```sh
# 本地编译
GOOS=linux GOARCH=amd64 go build
# 或？
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# docker 启动
docker-compose -f docker-compose-alpine.yml up -d
```

## TODO or 存在问题

- 返回数据结构待整理
  - 统一定义错误码
- 数据缓存，ES？redis？
  - 目前基于 orm 的 redis 数据缓存，连表查询不理想，待封装新的缓存机制
- 日志搜集，ES？
- 模型定义时为了插入 `null` 值，字段类型必须使用指针
  - 如果使用`database/sql`或`github.com/guregu/null`包，则导致 `validator.v8` 数据校验失效
  - 结构体转 json，忽略部分字段， tags 的方式貌似无效，待研究
  - 时间字段格式化，待研究
    - Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
- 基于 `binding` tags 的校验有些简陋，待优化
  - 只读字段校验，待研究
- 批量更新时 model 的 Hooks 不会运行
- 任务队列可视化管理问题
