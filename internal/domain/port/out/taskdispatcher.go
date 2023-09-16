package out



type TaskEventDispatcher interface {
	SendTaskFinishedEvent(int64) error
	SendTaskCanceledEvent(int64) error
}