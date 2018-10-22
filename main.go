package main

import (
	"wegou/config"
	"wegou/routes"
	"wegou/service/consumer"
)

func init() {
	config.InitLog()
}

func main() {

	//启动服务
	kafkaConfig := config.GetKafkaConfig()

	go (&consumer.FanConCumer{}).Consumer(kafkaConfig)
	go (&consumer.MaterialConsumer{}).Consumer(kafkaConfig)

	r := routes.Routes()
	webConfig := config.GetWebConfig()
	addr := webConfig.Host + ":" + webConfig.Port
	r.Run(addr)

}
