# gin-mall

**基于 gin+gorm+mysql读写分离 的一个电子商场**

# 项目运行
## 手动运行
普通运行
```go
cd ./cmd
go run main.go
```
以二进制文件运行
```go
go mod tidy
cd ./cmd
go build -o ../main
./main
```
**注意：手动运行方式不支持Skywalking，如果需要支持Skywalking，可以参考Makefile，按照文件内命令编译**
## 脚本运行
项目根目录内置了 Dockerfile、Makefile、docker-compose.yml 等文件
目的是快速构建项目环境，简易化项目运行难度

下面介绍 Makefile 中内置的几条指令，可根据需要在控制台**当前项目根目录下**进行相应操作的执行
```bash
make tools          # 构建Skywalking-Agent二进制文件
make                # 构建二进制文件并自动运行
make build          # 构建二进制文件
make env-up         # 拉起项目环境
make env-down       # 停止并删除环境
make docker-up      # 以容器方式拉起项目
make docker-down    # 停止并删除容器
```
对于第一次运行本项目，可以按如下顺序执行
1. 在Makefile中的前几行修改`ARCH`和`OS`以对应自己的电脑系统，在注释中提供了可选项
2. 运行如下代码
```bash
make env-up tools build # 拉起项目环境、编译Agent、构建项目二进制文件
./main                  # 运行项目
```

# 主要功能

- 用户注册登录(jwt-go)
- 用户基本信息修改，解绑定邮箱，修改密码
- 商品的发布，浏览等
- 购物车的加入，删除，浏览等
- 订单的创建，删除，支付等
- 地址的增加，删除，修改等
- 各个商品的浏览次数，以及部分种类商品的排行
- 设置了支付密码，对用户的金额进行了对称加密
- 支持事务，支付过程发送错误进行回退处理
- 可以将图片上传到对象存储，也可以切换分支上传到本地static目录下
- 添加ELK体系，方便日志查看和管理

# 项目结构
```
gin-mall
├── api             # 用于定义接口函数，也就是controller的作用
├── cmd             # 程序入口
├── conf            # 配置文件
├── doc             # 文档
├── middleware      # 中间件
├── model           # 数据库模型
├── pkg
│  ├── e            # 错误码
│  └── util         # 工具函数
├── repository
│  ├── cache        # Redis缓存
│  ├── db           # 持久层的mysql
│  │  ├── dao       # dao层，对db进行操作
│  │  └── model     # 定义mysql的模型
│  ├── es           # ElasticSearch，形成elk体系
│  └── mq           # 放置各种mq，kafka，rabbitmq等等
├── routes          # 路由逻辑处理
├── serializer      # 将数据序列化为 json 的函数，便于返回给前端
├── service         # 接口函数的实现
└── static          # 存放静态文件
```