package tasks

import (
	"errors"
	"fmt"
)

type Task struct {
	ID           int
	Name, Status string
}
type TaskList struct {
	tasks []Task
	n, s  int
}

func NewTask(id int, name, sta string) Task {

	return Task{id, name, sta}

}

func NewTaskList(size int) TaskList {

	return TaskList{make([]Task, 0, size), 0, size}

}

func (tasklist *TaskList) Create(task Task) error {

	if tasklist.n < tasklist.s {
		tasklist.tasks = append(tasklist.tasks, task)
		tasklist.n++
		return nil
	}
	return errors.New("Stopped. Length of TaskList Out of range!!")
}

func ReadTask(task *Task) string {

	return fmt.Sprintf("Id: %v, Name: %s, Status: %s", task.ID, task.Name, task.Status)

}

func (tasklist *TaskList) Read(index int) Task {

	return tasklist.tasks[index]

}

func (tasklist *TaskList) Update(index int, newtask Task) {

	if index < len(tasklist.tasks) {
		tasklist.tasks[index] = newtask
	}

}
func (taskList *TaskList) delete(index int) {
	
	if index < len(taskList.tasks) {
		taskList.tasks = append(taskList.tasks[:index],
			taskList.tasks[index+1:]...)
	}
}
/*
func main() {
	//t := Task{}
	tasklist := TaskList{}
	tasklist.Create(NewTask(150, "Bubble Sort", "Old"))
	tasklist.Create(NewTask(130, "Pascal triangle", "Almostly new"))
	tasklist.Create(NewTask(120, "Pifagor theory", "Old"))
	tasklist.Update(1, NewTask(110, "My theory", "Freshly new"))

	for i, tas := range tasklist.tasks {

		fmt.Printf("\n%v:  %v\n", i, tas)

	}
}
*/

