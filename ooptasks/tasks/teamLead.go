package tasks

import "fmt"

type TeamLead struct {
	Id   int
	Name string
}

func (temLead *TeamLead) String() string {
	return fmt.Sprintf("\n %v. Team Lead: %s\n", temLead.Id, temLead.Name)
}

var LeadId int

func NewLead(name string) TeamLead {
	temLead := TeamLead{LeadId, name}
	LeadId++
	return temLead
}

func (temLead *TeamLead) delegateDevTask(id int, step string) (string, error) {
	task, err := getTask(id)
	setTask(id, Task{task.taskId, task.name, step})
	task, err = getTask(id)

	if err == nil {
		fmt.Printf("%v ordered task : %v is %s\n", temLead.Name, id, task.status)
	}
	return task.status, err
}
