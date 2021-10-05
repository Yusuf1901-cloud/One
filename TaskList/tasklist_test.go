package tasks

import "testing"

const id = 31519
const name = "Mirmuhsin"
const status = "Resently Changed"
const n = 5

func TestNewTask(t *testing.T) {
	task := NewTask(id, name, status)

	if task.Name != name {
		t.Errorf("Expected: %s, found: %s\n", name, task.Name)
	}

	if task.ID != id {
		t.Errorf("Expected: %v, found: %v\n", id, task.ID)
	}

	if task.Status != status {
		t.Errorf("Expected: %s, found: %s\n", status, task.Status)
	}
}

func TestNewTaskList(t *testing.T) {
	taskList := NewTaskList(n)

	if taskList.n != n {
		t.Errorf("Found an error on taskList's length!!\n")
	}
}

func TestCreateTasklist(t *testing.T) {
	task := NewTask(id, name, status)
	taskList := NewTaskList(n)
	taskList.Create(task)
	if len(taskList.tasks) != 1 {

		t.Errorf("Expected: 1, founded : %v\n", len(taskList.tasks))

	}

	if taskList.tasks[0] != task {

		t.Errorf("Founded: %v, not found: %v\n", taskList.tasks[0], task)

	}
}

func TestRead(t *testing.T) {
	task := NewTask(id, name, status)
	taskList := NewTaskList(n)

	for i := 0; i < n; i++ {
		taskList.Create(task)
	}

	if taskList.tasks[0] != task {
		t.Errorf("Expected: %v, found: %v\n", task, taskList.tasks[0])
	}

}

func TestUpdateTasks(t *testing.T) {
	task := NewTask(id, name, status)
	taskList := NewTaskList(n)

	taskcha := NewTask(32019, "Fibonacci", "Recently solved!!")

	for i := 0; i < n; i++ {
		taskList.Create(task)
	}

	taskList.Update(1, taskcha)

	if taskList.tasks[1] != taskcha {

		t.Errorf("Expected: %v, found: %v\n", taskcha, taskList.tasks[1])

	}
}

func TestDeleteTask(t *testing.T) {
	task := NewTask(id, name, status)
	taskList := NewTaskList(n)

	for i := 0; i < n; i++ {
		taskList.Create(task)
	}

	taskcha := NewTask(32019, "Fibonacci", "Recently solved!!")
	taskList.Update(1, taskcha)

	taskList.delete(1)
	if taskList.tasks[1] == taskcha {

		t.Errorf("\nError occured, First element of tasks should be: %v, but found: %v\n", task, taskcha)

	}

}
