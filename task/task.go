package task

import "taskTracker/utils"

const (
	TODO  = 0
	DOING = 1
	DONE  = 2
)

type Task struct {
	id          int
	discription string
	status      int
}

func LoadTask(id int, des string, status int) *Task {
	return &Task{
		id:          id,
		discription: des,
		status:      status,
	}
}

func NewTask(discription string) *Task {
	return &Task{
		id:          utils.GenerateID(),
		discription: discription,
		status:      TODO,
	}
}
func (t *Task) Todo() {
	t.status = TODO
}

func (t *Task) Doing() {
	t.status = DOING
}

func (t *Task) Done() {
	t.status = DONE
}

func (t *Task) ID() int {
	return t.id
}

func (t *Task) Describe() string {
	return t.discription
}

func (t *Task) Status() int {
	return t.status
}

func (t *Task) Update(discription string) {
	t.discription = discription
}
