package task

import (
	"strings"

	"github.com/spf13/viper"
)

type Kafka struct {
	Blockers       []string
	CustomerTopics []string
	MaterialTopics []string
}

func GetKafkaConfig() Kafka {
	conf := Kafka{
		Blockers:       strings.Split(viper.GetString("kafka.broker"), ","),
		CustomerTopics: strings.Split(viper.GetString("kafak.customer_topic"), ","),
		MaterialTopics: strings.Split(viper.GetString("kafak.material_topic"), ","),
	}
	return conf
}
