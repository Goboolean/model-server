package manager

import (
	"context"
	"time"

	"github.com/Goboolean/model-server/internal/domain/port/in"
	"github.com/Goboolean/model-server/internal/domain/port/out"
	process "github.com/Goboolean/model-server/internal/domain/service/processset"
	task "github.com/Goboolean/model-server/internal/domain/service/taskqueue"
	worker "github.com/Goboolean/model-server/internal/domain/service/workerqueue"
)



type Manager struct {
	t *task.Queue
	w *worker.Queue
	p *process.Set

	td out.TaskEventDispatcher

	ctx context.Context
	cancel context.CancelFunc
}

func New(t *task.Queue, w *worker.Queue, p *process.Set) (in.EventHandler, error) {

	ctx, cancel := context.WithCancel(context.Background())

	instance := &Manager{
		t: t,
		w: w,
		p: p,

		ctx: ctx,
		cancel: cancel,
	}

	//instance.run()

	return instance, nil
}



func (m *Manager) run(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(time.Second):
				if m.t.IsEmpty() || m.w.IsEmpty() {
					continue
				}

				t, err := m.t.Dequeue()
				if err != nil {
					continue
				}

				w, err := m.w.Dequeue()
				if err != nil {
					m.t.Enqueue(t)
					continue
				}

				if err := m.p.Insert(t, w); err != nil {
					m.t.Enqueue(t)
					m.w.Enqueue(w)
					continue
				}
			}
		}
	}(ctx)
}




