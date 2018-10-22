package consumer

import (
	"os"
	"os/signal"
	"wegou/config"
	"wegou/service/task"

	"github.com/bsm/sarama-cluster"
)

type MaterialConsumer struct{}

//素材任务消费者
func (materialConsumer *MaterialConsumer) Consumer(kafkaConfig config.Kafka) {

	// init (custom) config, set mode to ConsumerModePartitions
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions

	// init consumer
	brokers := kafkaConfig.Blockers
	topics := kafkaConfig.MaterialTopics
	consumer, err := cluster.NewConsumer(brokers, kafkaConfig.MaterialGroup, topics, config)
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
					dto := task.MaterialDto{}
					dto.Material(string(msg.Value))
					//fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
					consumer.MarkOffset(msg, "") // mark message as processed
				}
			}(part)
		case <-signals:
			return
		}
	}

}
