package main

import (
	"fmt"
)

type TaskList struct {
	ID int
	Name, Status, Priority, CreatedAt, CreatedBy, DueDate string
}

func ( cre *TaskList) Create(i int, a, b, c, d, e, f string){
	cre.ID = i
	cre.Name = a
	cre.Status = b
	cre.Priority = c
	cre.CreatedAt = d
	cre.CreatedBy = e
	cre.DueDate = f
}

func updateIdDueDate(up *TaskList, id int, date string){
	up.ID = id
	up.DueDate = date
}

func GetAll( get *TaskList) {
	fmt.Println("Task Id: ", get.ID)
	fmt.Printf("Your Task %s %s and %s\n", get.Name, get.Status, get.Priority)
	fmt.Printf("This task created by %s at %s\n", get.CreatedAt, get.CreatedBy)
	fmt.Printf("DueDate:%s\n\n", get.DueDate)
}

func main() {
	objects := []TaskList{
		{
			ID :546897,
			Name: "Pascal Triangle",
			Status: "New",
			Priority: "Required",
			CreatedAt: "13-01-2021",
			CreatedBy: "Ahmadbek",
			DueDate: "13-12-2021",
		},
		{
			ID :483197,
			Name: "Geron formula",
			Status: "Old",
			Priority: "Sometimes Required",
			CreatedAt: "1800-1900",
			CreatedBy: "Geron",
			DueDate: "13-12-2021",
		},
	}
	obj := TaskList{}
	obj.Create(
		201369, 
		"Pifagor theory", 
		"Very Old", 
		"Required", 
		"200-300 year",
		"Pifagor",
		"Unlimited",
	)
	objects = append(objects, obj)
	
	updateIdDueDate(&objects[0], 517201, "01-11-2021")
	
	for i, obj0 := range objects{
		fmt.Printf("\n\n%v-object:\n", i+1)
		GetAll(&obj0)
	}
	
	objects = []TaskList{} //deleting objects array
}
