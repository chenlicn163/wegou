package task

import (
	"fmt"
	"os"
	"time"
	"wegou/config"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
)

type Task struct{ Topics string }

//异步提交任务
func (task *Task) AsyncProducer(value string) {
	kafkaConfig := config.GetKafkaConfig()
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewAsyncProducer(kafkaConfig.Blockers, config)
	defer p.Close()
	if err != nil {
		return
	}

	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					glog.Errorln(err)
				}
			case <-success:
			}
		}
	}(p)

	fmt.Fprintln(os.Stdout, value)
	msg := &sarama.ProducerMessage{
		Topic: task.Topics,
		Value: sarama.ByteEncoder(value),
	}
	p.Input() <- msg
}
