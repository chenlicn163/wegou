package engine

import "wegou/service/task"

func Progress() {
	go task.CustomerConsumer()
	go task.MaterialConsumer()
}
