package task

type StatusTate interface {
	Todo() string
	Doing() string
	Done() string
}

type Status struct {
	state StatusTate
}

type Todo struct {
	task *Task
}

func (t *Todo) Todo() string {
	return "Todo"
}

func (t *Todo) Doing() string {
	t.task.Doing()
	return "Doing"
}

func (t *Todo) Done() string {
	t.task.Done()
	return "Done"
}

type Doing struct {
	task *Task
}

func (d *Doing) Todo() string {
	d.task.Todo()
	return "Todo"
}

func (d *Doing) Doing() string {
	return "Doing"
}

func (d *Doing) Done() string {
	d.task.Done()
	return "Done"
}

type Done struct {
	task *Task
}

func (d *Done) Todo() string {
	d.task.Todo()
	return "Todo"
}

func (d *Done) Doing() string {
	d.task.Doing()
	return "Doing"
}
func (d *Done) Done() string {
	return "Done"
}

func NewStatus(task *Task) *Status {
	status := &Status{}
	switch task.Status() {
	case TODO:
		status.state = &Todo{task: task}
	case DOING:
		status.state = &Doing{task: task}
	case DONE:
		status.state = &Done{task: task}
	}
	return status
}
