package tasks

import (
	"errors"
)

type Task struct {
	tasdId       int
	name, status string
}

var tasks map[int]Task
var taskID int

func NewTask(name string) Task {
	newTask := Task{taskID, name, "todo"}

	if tasks == nil {
		tasks = make(map[int]Task)
	}
	tasks[taskID] = newTask
	taskID++
	return newTask
}

func getTask(id int) (Task, error) {
	task, ok := tasks[id]
	err := errors.New("Task not found")
	if ok {
		err = nil
	}
	return task, err
}

func setTask(id int, task Task) {

	tasks[id] = task

}
