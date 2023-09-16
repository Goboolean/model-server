package task

import (
	"sync"

	"github.com/Goboolean/model-server/internal/domain/entity"

	"container/list"
)


const taskptr = "taskptr"



type Queue struct {
	mu sync.Mutex

	list list.List
}


func New() *Queue {
	return &Queue{
		list: *list.New(),
	}
}



func (q *Queue) Enqueue(t *entity.Task) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushBack(t)

	return nil
}


func (q *Queue) EnqueueLeft(t *entity.Task) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.list.PushFront(t)

	return nil
}


func (q *Queue) Dequeue() (*entity.Task, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Length() == 0 {
		return nil, ErrQueueEmpty
	}

	v := q.list.Remove(q.list.Front())
	w := v.(*entity.Task)	
	return w, nil
}


func (q *Queue) Remove(id int64) (*entity.Task, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for e := q.list.Front(); e != nil; e = e.Next() {
		w := e.Value.(*entity.Task)
		if w.ID == id {
			q.list.Remove(e)
			return w, nil
		}
	}

	return nil, ErrTaskNotFound
}


func (q *Queue) Exists(id int64) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for e := q.list.Front(); e != nil; e = e.Next() {
		w := e.Value.(*entity.Task)
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