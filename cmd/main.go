package main

import (
	"fmt"

	conf "github.com/CocaineCong/gin-mall/config"
	util "github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/pkg/utils/track"
	"github.com/CocaineCong/gin-mall/repository/cache"
	"github.com/CocaineCong/gin-mall/repository/db/dao"
	"github.com/CocaineCong/gin-mall/repository/es"
	"github.com/CocaineCong/gin-mall/repository/kafka"
	"github.com/CocaineCong/gin-mall/repository/rabbitmq"
	"github.com/CocaineCong/gin-mall/routes"

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
