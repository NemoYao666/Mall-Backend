package main

import (
	"fmt"

	conf "gin-mall-backend/config"
	util "gin-mall-backend/pkg/utils/log"
	"gin-mall-backend/pkg/utils/track"
	"gin-mall-backend/repository/cache"
	"gin-mall-backend/repository/db/dao"
	"gin-mall-backend/repository/es"
	"gin-mall-backend/repository/kafka"
	"gin-mall-backend/repository/rabbitmq"
	"gin-mall-backend/routes"

	_ "github.com/apache/skywalking-go"
)

func main() {
	loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}

// loading一些配置
func loading() {
	conf.InitConfig()       // viper读取配置文件
	dao.InitMySQL()         // 数据库主从架构读写分离
	cache.InitCache()       // redis 初始化
	rabbitmq.InitRabbitMQ() // RabbitMQ 初始化
	es.InitEs()             // ELK 初始化
	kafka.InitKafka()       // kafka 初始化
	track.InitJaeger()      // jaeger 初始化
	util.InitLog()          // 接入ELK日志初始化
	fmt.Println("加载配置完成...")
	go scriptStarting()
}

func scriptStarting() {
	// 启动一些脚本
}
