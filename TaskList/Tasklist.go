package main

import "fmt"

type Task struct {
	ID           int
	Name, Status string
}
type TaskList struct {
	tasks []Task
	n     int
}

func (task *Task) NewTask(id int, name, sta string) Task {
	task.ID = id
	task.Name = name
	task.Status = sta

	return *task
}

func ReadTask(task *Task) string {
	return fmt.Sprintf("Id: %v, Name: %s, Status: %s", task.ID, task.Name, task.Status)
}

func (tasklist *TaskList) Create(task Task) {
	tasklist.tasks = append(tasklist.tasks, task)
}

func (tasklist *TaskList) ReadAll(index int) Task {

	return tasklist.tasks[index]

}

func (tasklist *TaskList) Update(index int, newtask Task) {

	if index < len(tasklist.tasks) {
		tasklist.tasks[index] = newtask
	}

}

func main() {
	t := Task{}
	tasklist := TaskList{}
	tasklist.Create(t.NewTask(150, "Bubble Sort", "Old"))
	tasklist.Create(t.NewTask(130, "Pascal triangle", "Almostly new"))
	tasklist.Create(t.NewTask(120, "Pifagor theory", "Old"))
	tasklist.Update(1, t.NewTask(110, "My theory", "Freshly new"))

	for i, tas := range tasklist.tasks {

		fmt.Printf("\n%v:  %v\n", i, tas)

	}
}
