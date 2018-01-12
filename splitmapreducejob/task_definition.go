package splitmapreducejob

// TaskDefinition is the canonical computes
// task definition format
type TaskDefinition struct {
	Runner *Runner               `json:"runner"`
	Result *TaskDefinitionResult `json:"result"`
}

// TaskDefinitionResult defines what to do with the
// result after the task is done
type TaskDefinitionResult struct {
	Action      string                           `json:"action"`
	Destination *TaskDefinitionResultDestination `json:"destination"`
}

// TaskDefinitionResultDestination describes a place
// a result can go
type TaskDefinitionResultDestination struct {
	Dataset *Location `json:"dataset"`
	Path    string    `json:"path"`
}
