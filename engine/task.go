package engine

import "wegou/engine/task"

func Progress() {
	go task.CustomerConsumer()
	go task.MaterialConsumer()
}
