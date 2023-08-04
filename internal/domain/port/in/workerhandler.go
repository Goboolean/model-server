package port



type WorkerHandler interface {
	RegisterWorker() error
	ReportWorkerConnectionLost(workerId int64) error
}