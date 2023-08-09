package manager

import "github.com/Goboolean/model-server/internal/domain/entity"




func (m *Manager) RegisterTask(t *entity.Task) error {

	if err := m.t.Enqueue(t); err != nil {
		return err
	}

	// TODO : Check if enable to create worker task pair
	return nil
}


func (m *Manager) CancelTask(id int64) error {

	// Check if task exists on task queue
	if exist := m.t.Exists(id); exist {
		_, err := m.t.Remove(id)
		if err != nil {
			return err
		}
		return nil
	}

	// Check if task exists on process set
	if exist := m.p.Exists(id); exist {

		process, err := m.p.Remove(id)
		if err != nil {
			return err
		}

		if err := process.Worker().Cancel(); err != nil {
			return err
		}

		if err := m.td.SendTaskCanceledEvent(id); err != nil {
			return err
		}

		return nil
	}

	// Unown Task ID
	return ErrUnknownTaskID
}