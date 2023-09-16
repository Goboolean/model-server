package entity



type Process struct {
	w *Worker
	t *Task
}

func NewProcess(w *Worker, t *Task) *Process {
	return &Process{w: w, t: t}
}

func (p *Process) Worker() *Worker {
	return p.w
}

func (p *Process) Task() *Task {
	return p.t
}