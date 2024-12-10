package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"taskTracker/task"
	"taskTracker/utils"
)

type Manager struct {
	tasks map[int]*task.Task
	done  []*task.Task
	todo  []*task.Task
	doing []*task.Task
}

type FileTask struct {
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	Status int    `json:"status"`
}

type FileTasks struct {
	ID    int         `json:"id"`
	Tasks []*FileTask `json:"tasks"`
}

func NewManager(cap int) *Manager {
	return &Manager{
		tasks: make(map[int]*task.Task, cap),
		done:  make([]*task.Task, 0),
		todo:  make([]*task.Task, 0),
		doing: make([]*task.Task, 0),
	}
}

func (m *Manager) LoadTasks(r io.Reader) {
	FileTasks := FileTasks{Tasks: make([]*FileTask, 0)}
	json.NewDecoder(r).Decode(&FileTasks)
	for _, ft := range FileTasks.Tasks {
		if _, ok := m.tasks[ft.ID]; ok {
			continue
		}
		t := task.LoadTask(ft.ID, ft.Desc, ft.Status)
		m.tasks[t.ID()] = t
		if t.Status() == task.DONE {
			m.done = append(m.done, t)
		}
		if t.Status() == task.TODO {
			m.todo = append(m.todo, t)
		}
		if t.Status() == task.DOING {
			m.doing = append(m.doing, t)
		}
	}
	utils.SetID(FileTasks.ID)
}

func (m *Manager) SaveTasks(w io.Writer) {

	FileTasks := FileTasks{Tasks: make([]*FileTask, 0)}
	for _, t := range m.tasks {
		ft := FileTask{ID: t.ID(), Desc: t.Describe(), Status: t.Status()}
		FileTasks.Tasks = append(FileTasks.Tasks, &ft)
	}
	FileTasks.ID = utils.CurID()

	json.NewEncoder(w).Encode(FileTasks)

}

func (m *Manager) AddTask(t *task.Task) {
	m.tasks[t.ID()] = t
	m.todo = append(m.todo, t)
}

func (m *Manager) GetTask(id int) *task.Task {
	if _, ok := m.tasks[id]; !ok {
		return nil
	}
	return m.tasks[id]
}

func (m *Manager) DeleteTask(id int) {
	t := m.GetTask(id)
	if t == nil {
		return
	}
	if t.Status() == task.DOING {
		removeTask(m.doing, t)
	}
	if t.Status() == task.TODO {
		removeTask(m.todo, t)
	}
	if t.Status() == task.DONE {
		removeTask(m.done, t)
	}
	delete(m.tasks, id)
}

func (m *Manager) UpdateTask(id int, to string) error {
	task := m.GetTask(id)
	if task == nil {
		return errors.New("task not found")
	}
	task.Update(to)
	return nil
}

func (m *Manager) Doing(id int) error {
	t := m.GetTask(id)
	if t == nil {
		return errors.New("task not found")
	}
	if t.Status() == task.TODO {
		m.todo2doing(t)
		t.Doing()
		return nil
	}
	return nil
}

func (m *Manager) Done(id int) error {
	t := m.GetTask(id)
	if t == nil {
		return errors.New("task not found")
	}
	if t.Status() == task.DOING {
		m.doing2done(t)
		t.Done()
		return nil
	}
	if t.Status() == task.TODO {
		m.todo2done(t)
		t.Done()
		return nil
	}
	return nil
}
func (m *Manager) ListAll() {
	for _, t := range m.tasks {
		fmt.Printf("%d: %s \n", t.ID(), t.Describe())
	}
}

func (m *Manager) ListTodo() {
	for _, t := range m.todo {
		fmt.Printf("%d: %s \n", t.ID(), t.Describe())
	}
}

func (m *Manager) ListDoing() {
	for _, t := range m.doing {
		fmt.Printf("%d: %s \n", t.ID(), t.Describe())
	}
}

func (m *Manager) ListDone() {
	for _, t := range m.done {
		fmt.Printf("%d: %s \n", t.ID(), t.Describe())
	}
}

func (m *Manager) doing2done(t *task.Task) {
	m.doing = removeTask(m.doing, t)
	m.done = append(m.done, t)
}

func (m *Manager) todo2doing(t *task.Task) {
	m.todo = removeTask(m.todo, t)
	m.doing = append(m.doing, t)
}

func (m *Manager) todo2done(t *task.Task) {
	m.todo = removeTask(m.todo, t)
	m.done = append(m.done, t)
}

func removeTask(tasks []*task.Task, t *task.Task) []*task.Task {
	for i, task := range tasks {
		if task.ID() == t.ID() {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	return tasks
}
