package main

import (
	"fmt"
)

type ContactList struct {
	ID, Phone int
	FirstName, LastName, Email, Position string
	
}

func (y *ContactList) Create(ph,id int, a, b, c, d string) {
	y.ID = ph
	y.Phone = id
	y.FirstName = a
	y.LastName = b
	y.Email = c
	y.Position = d
	fmt.Println("Objects created!\n")
}

func getIDPhone(tel *ContactList) {
	
	fmt.Printf("Your id: %v with %v phone number\n", tel.ID, tel.Phone)
	
}

func (up *ContactList) updateIdPhone(id, phone int){
	up.ID = id
	up.Phone = phone
}


func (g *ContactList) getAll(){
	fmt.Println("Your ID : ", g.ID)
	fmt.Printf("Your phone: %v\n",  g.Phone)
	fmt.Printf("You are %s  %s\n", g.FirstName, g.LastName)
	fmt.Println("Your email:", g.Email)
	fmt.Printf("Now you are working at UDEV company as %s position", g.Position)
}

func main() {
	y := ContactList{} //fields : ID, Phone, Name, LastName, Email, Position
	y.Create(31019, 993654858, "Xalil", "Jalilov", "xjalilov@gmail.com", "Junior developer")
	
	getIDPhone(&y)
	
	y.updateIdPhone(31419,901471901)
	
	fmt.Println("After Updated ID and Phone: \n")
	
	y.getAll()
	
	y = ContactList{} //delete all objects
	
}
