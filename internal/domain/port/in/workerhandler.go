package in

import "github.com/Goboolean/model-server/internal/domain/port/out"



type WorkerRegisterer interface {
	RegisterWorker(out.WorkerSession) error
	UnregisterWorker(int64) error
}

type WorkerEventHandler interface {
	HandleWorkerConnectionLost(int64) error
	HandleWorkerConnectionRestored(int64) error
	HandleWorkerFinished(int64) error
}