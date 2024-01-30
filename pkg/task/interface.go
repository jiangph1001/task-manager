package task

type TaskStatus int

const (
	Created TaskStatus = iota
	Running
	Stopped
	Completed
	Failed
)

// Task is an interface that all tasks must implement
type Task interface {
	Execute() error
	Status() TaskStatus
	SetStatus(status TaskStatus)
}
