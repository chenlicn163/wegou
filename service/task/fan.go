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

type FanConCumer struct{}

//粉丝任务消费者
func (fanConCumer *FanConCumer) Consumer(kafkaConfig config.Kafka) {

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
					fanConCumer.customer(msg)
					consumer.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		case <-signals:
			return
		}
	}
}

func (fanConCumer *FanConCumer) customer(msg *sarama.ConsumerMessage) {
	message := string(msg.Value)
	topic := gjson.Get(message, "kafka.topic").String()
	switch topic {
	case "customer-add":
		fanConCumer.customerAdd(message)
	}
}

func (fanConCumer *FanConCumer) customerAdd(message string) {
	web := gjson.Get(message, "web").String()
	openid := gjson.Get(message, "openid").String()
	fan := wx.GetFan(web, openid).Get()
	fan.AddFan(web)
}
