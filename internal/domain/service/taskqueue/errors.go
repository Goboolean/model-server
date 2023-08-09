package task



import "errors"

var ErrQueueEmpty = errors.New("queue is empty")

var ErrTaskNotFound = errors.New("worker is not found")