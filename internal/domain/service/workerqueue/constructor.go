package worker

import (
	"container/list"
	"context"
	"sync"

	"github.com/Goboolean/model-server/internal/domain/entity"
)


const workerptr = "workerptr"



type Queue struct {
	mu sync.Mutex

	list list.List

	canceledChan chan *entity.Worker
	

	ctx context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}


func New() *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		ctx: ctx,
		cancel: cancel,

		canceledChan: make(chan *entity.Worker),
	}
}


func (q *Queue) Enqueue(w *entity.Worker) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushBack(w)
	return nil
}


func (q *Queue) EnqueueLeft(t *entity.Task) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushFront(t)

	return nil
}

func (q *Queue) Dequeue() (*entity.Worker, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	v := q.list.Remove(q.list.Front()).(*entity.Worker)
	return v, nil
}


func (q *Queue) Remove(id int64) (*entity.Worker, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for e := q.list.Front(); e != nil; e = e.Next() {
		w := e.Value.(*entity.Task)
		if w.ID == id {
			q.list.Remove(e)
		}
	}

	return nil, ErrWorkerNotFound
}


func (q *Queue) Exists(id int64) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for e := q.list.Front(); e != nil; e = e.Next() {
		w := e.Value.(*entity.Worker)
		if w.ID == id {
			return true
		}
	}

	return false
}

func (q *Queue) Length() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.list.Len() == 0
}



func (q *Queue) Close() {
	q.cancel()
	q.wg.Wait()
}