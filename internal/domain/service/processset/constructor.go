package process

import (
	"sync"

	"github.com/Goboolean/model-server/internal/domain/entity"
)




type Set struct {
	mu sync.RWMutex

	set map[int64]*entity.Process
}

type Atom struct {
	w *entity.Worker
	t *entity.Task
}


func New() *Set {
	return &Set{
		set: make(map[int64]*entity.Process),
	}
}

func (s *Set) Insert(t *entity.Task, w *entity.Worker) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := t.ID
	if _, ok := s.set[id]; ok {
		return ErrAtomAlreadyExists
	}
	s.set[id] = entity.NewProcess(w,t)

	return nil
}


func (s *Set) Remove(id int64) (*entity.Process, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	atom, ok := s.set[id];
	if !ok {
		return nil, ErrAtomNotFound
	}

	delete(s.set, id)
	return atom, nil
}


func (s *Set) Exists(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.set[id]; ok {
		return true
	}
	return false
}