# For unix-based system, such as macOS/Linux, ARCH can be replace with $(shell uname -m)
# For windows, can use 'systeminfo | findstr /C:"System Type"' to get ARCH
# Or easily, just use 'amd64' for windows

# amd64 arm64
ARCH = amd64

# linux darwin windows
OS = linux

DIR := $(shell pwd)
OUTPUT = main

CONTAINER_NAME = gin_mall_server
IMAGE_NAME = gin_mall:3.0

GO = go
GO_BUILD = $(GO) build
GO_BUILD_FLAGS = -v

.PHONY: run			# 构建同时运行
test:
	@make build
	@./$(OUTPUT)

.PHONY: build		# 构建项目
build:
	@echo "build project to ./$(OUTPUT)"
	$(GO_BUILD) \
	-a -o ./$(OUTPUT) ./cmd

.PHONY: env-up		# 启动环境
env-up:
	docker-compose up -d
	@echo "env start success"

.PHONY: env-down	# 关闭环境
env-down:
	docker-compose down
	@echo "env stop success"

.PHONY: docker-up	# 以容器形式部署项目
docker-up:
	docker build \
	-t $(IMAGE_NAME) \
	-f ./Dockerfile \
	./
	docker run \
	-it \
	--name $(CONTAINER_NAME) \
	--network host \
	-d $(IMAGE_NAME)
	@echo "container run success at localhost:5001"

.PHONY: docker-down # 结束docker部署,同时删除容器和镜像
docker-down:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)
	@echo "container stop && rm success"

default: run