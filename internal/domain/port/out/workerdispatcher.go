package out



type WorkerSession interface {
	AllocateTask(id string, model string, stockId string) error
	CancelWOrk() error
	ClosedEvent() chan struct{}
	FinishedEvent() chan struct{}
}