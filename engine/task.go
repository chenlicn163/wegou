package engine

import "wegou/engine/task"

func Progress() {
	kafkaConfig := GetKafkaConfig()
	go task.CustomerConsumer(kafkaConfig)
	go task.MaterialConsumer(kafkaConfig)
}
