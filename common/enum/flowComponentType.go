package enum

type FlowComponentType string

const (
	Trigger   FlowComponentType = "TRIGGER"
	Function  FlowComponentType = "FUNCTION"
	Output    FlowComponentType = "OUTPUT"
	Scheduler FlowComponentType = "SCHEDULER"
)
