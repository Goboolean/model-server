package in



type EventHandler interface {
	TaskHandler
	WorkerRegisterer
	WorkerEventHandler
}