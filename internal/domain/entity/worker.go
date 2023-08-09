package entity

import (
	"context"

	"github.com/Goboolean/model-server/internal/domain/port/out"
)



type WorkerStatus int

const (
	WorkerStatusIdle WorkerStatus = iota + 1
	WorkerStatusWorking
)

type WorkerSession interface {
	
}


var (
	id int64
)

type Worker struct {
	ID int64
	session out.WorkerSession

	ctx context.Context
	cancel context.CancelFunc
}

func NewWorker(session out.WorkerSession) *Worker {
	id++
	return &Worker{
		ID: id,
		session: session,
	}
}

func (w *Worker) Context() context.Context {
	return w.ctx
}

func (w *Worker) Cancel() error {
	w.cancel()
	return nil
}