package tasks

import (
	"fmt"
)

type Developer struct {
	Id   int
	Name string
}

func (delevoper *Developer) String() string {
	return fmt.Sprintf("\n %v. Developer: %s\n", delevoper.Id, delevoper.Name)
}

var DevId int

func NewDev(name string) Developer {
	developer := Developer{DevId, name}
	DevId++
	return developer
}

func (developer *Developer) develop(id int, step string) error {
	task, err := getTask(id)

	if err == nil {
		fmt.Printf("%v developed: %v\n", developer, task.name)
		task.status = "done"
		setTask(id, task)
	}
	return err
}
