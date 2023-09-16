package entity




type TaskStatus int

const (
	TaskStatusIdle TaskStatus = iota + 1
	TaskStatusWorking
	TaskStatusDone
	TaskStatusCancelled
	TaskStatusFailed
)


type TaskSession interface {

}



type Task struct {
	ID int64

	Model string
	StockId string
}

func NewTask(id int64) *Task {
	return &Task{ID: id}
}


func (t *Task) GetID() int64 {
	return 0
}

func (t *Task) GetStatus() TaskStatus {
	return TaskStatus(0)
}

func (t *Task) SetStatusWorking() error {
	return nil
}

func (t *Task) SetStatusDone() error {
	return nil
}

