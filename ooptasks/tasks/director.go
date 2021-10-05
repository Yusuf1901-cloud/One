package tasks

import "fmt"

type Director struct {
	Id   int
	Name string
}

func (director *Director) String() string {
	return fmt.Sprintf("\n %v. Director: %s\n", director.Id, director.Name)
}

var directorId int

func NewDirector(name string) Director {
	director := Director{directorId, name}
	directorId++
	return director
}

func (direct *Director) giveTask(tasName string) (int, error) {
	task := NewTask(tasName)
	fmt.Printf("%s gave task %v\n", direct.Name, task)
	return task.taskId, nil

}
