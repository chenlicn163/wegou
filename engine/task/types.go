package task

type Kafka struct {
	Blockers       []string
	CustomerTopics []string
	MaterialTopics []string
	CustomerGroup  string
	MaterialGroup  string
}
