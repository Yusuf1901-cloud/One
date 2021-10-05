package main

import (
	"tasks/tasks"
)

func main() {
	var developer tasks.Developer = tasks.NewDev("Ahmadbek")
	director := tasks.NewDirector("Yusufbek")
	teamLead := tasks.NewLead("Amirbek")

	task, _ := director.giveTask("something")
	teamLead.Delegate(task, "initial")
	developer.develop(task, "dev")
	teamLead.Delegate(task, "done")

}
