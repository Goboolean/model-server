package port


type TaskHandler interface {
	RegisterTask() error
	CancelTask(taskid int64) error
}