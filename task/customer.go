package task

import (
	"fmt"
	"os"
	"os/signal"
	"wegou/types"

	"github.com/tidwall/gjson"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func CustomerConsumer(kafkaConfig types.Kafka) {

	// init (custom) config, set mode to ConsumerModePartitions
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions

	// init consumer
	brokers := kafkaConfig.Blockers
	topics := kafkaConfig.CustomerTopics
	consumer, err := cluster.NewConsumer(brokers, kafkaConfig.CustomerGroup, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume partitions
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				return
			}

			// start a separate goroutine to consume messages
			go func(pc cluster.PartitionConsumer) {
				for msg := range pc.Messages() {
					fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
					consumer.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		case <-signals:
			return
		}
	}
}

func customer(msg *sarama.ConsumerMessage) {
	message := string(msg.Value)
	topic := gjson.Get(message, "kafka.topic").String()
	switch topic {
	case "customer-add":
		customerAdd(message)
	}
}

func customerAdd(message string) {

}
