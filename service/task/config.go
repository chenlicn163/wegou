package task

import (
	"strings"

	"github.com/spf13/viper"
)

type Kafka struct {
	Blockers       []string
	CustomerTopics []string
	MaterialTopics []string
	CustomerGroup  string
	MaterialGroup  string
}

func GetKafkaConfig() Kafka {
	conf := Kafka{
		Blockers:       strings.Split(viper.GetString("kafka.broker"), ","),
		CustomerTopics: strings.Split(viper.GetString("kafak.customer_topic"), ","),
		MaterialTopics: strings.Split(viper.GetString("kafak.material_topic"), ","),
		CustomerGroup:  viper.GetString("kafka.customer_group"),
		MaterialGroup:  viper.GetString("kafka.material_group"),
	}
	return conf
}
