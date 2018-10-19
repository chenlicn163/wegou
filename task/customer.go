package task

import (
	"os"
	"os/signal"
	"wegou/config"
	"wegou/service/wx"

	"github.com/tidwall/gjson"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
)

type CustomerConCumer struct{}

func (customerConCumer *CustomerConCumer) Consumer(kafkaConfig config.Kafka) {

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
					/*fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
					 */
					customerConCumer.customer(msg)
					consumer.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		case <-signals:
			return
		}
	}
}

func (customerConCumer *CustomerConCumer) customer(msg *sarama.ConsumerMessage) {
	message := string(msg.Value)
	topic := gjson.Get(message, "kafka.topic").String()
	switch topic {
	case "customer-add":
		customerConCumer.customerAdd(message)
	}
}

func (customerConCumer *CustomerConCumer) customerAdd(message string) {
	web := gjson.Get(message, "web").String()
	openid := gjson.Get(message, "openid").String()
	fan := wx.GetCustomer(web, openid).Get()
	fan.AddFan(web)
}
