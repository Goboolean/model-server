package worker

import "errors"

var ErrQueueEmpty = errors.New("queue is empty")

var ErrWorkerNotFound = errors.New("worker is not found")