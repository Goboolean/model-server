package manager

import (
	"github.com/Goboolean/model-server/internal/domain/entity"
	"github.com/Goboolean/model-server/internal/domain/port/out"
)



func (m *Manager) RegisterWorker(sess out.WorkerSession) error {
	w := entity.NewWorker(sess)

	if err := m.w.Enqueue(w); err != nil {
		return err		
	}

	// TODO : Check if enable to create worker task pair
	return nil
}

func (m *Manager) UnregisterWorker(id int64) error {

	// Check if worker exists on task queue
	if exist := m.w.Exists(id); exist {
		_, err := m.w.Remove(id)
		if err != nil {
			return err
		}
		return nil
	}

	// Check if worker exists on process set
	if exist := m.p.Exists(id); exist {
	}

	// Unown Worker ID
	return ErrUnknownWorkerID
}


func (m *Manager) HandleWorkerConnectionLost(id int64) error {
	if exists := m.w.Exists(id); exists {
		_, err := m.w.Remove(id)
		return err
	}

	if exists := m.p.Exists(id); exists {
		_, err := m.p.Remove(id)
		return err
	}
	return ErrUnknownWorkerID

	// TODO: Change logic as only temporary connection lost not affect the process
}
func (m *Manager) HandleWorkerConnectionRestored(id int64) error {
	// TODO: Change logic as only temporary connection lost not affect the process
	if exists := m.w.Exists(id); exists {
		return nil
	}
	if exists := m.p.Exists(id); exists {
		return nil
	}
	return ErrUnknownWorkerID
}

func (m *Manager) HandleWorkerFinished(id int64) error {
	if exists := m.p.Exists(id); !exists {
		return ErrUnknownWorkerID
	}

	_, err := m.p.Remove(id)
	return err
}