package main

import (
	"contacts/contacts"
	"fmt"
)

const (
	hostname = "postgresql"
	hostPort = 5432
	user     = "postgres"
	password = "12345"
	dbName   = "postgres"
)
const n = 1

func main() {
	err := contacts.InitDB(hostname, hostPort, user, password, dbName)

	if err != nil {
		panic(err)
	}
	defer contacts.CloseDB()

	fmt.Println("Creating contacts:..")

	contact, err := contacts.NewContact(901235474, "komil")
	if err != nil {
		fmt.Printf("Error occured: %v.\n", err)
	}
	fmt.Println(contact)

	fmt.Println("Updating Contact")
	contact = contacts.UpdateContact(contact.ID, 952147475, "New name: Lobaroy")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(contact)

	fmt.Println("Getting all contacts: ")
	result, err := contacts.GetAll()

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	for _, contact := range result {
		fmt.Println(contact)
	}

	fmt.Println("Deleting contact:")
	err = contacts.DeleteContact(contact.getId())
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	result, err = contacts.GetAll()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	for _, contact := range result {
		fmt.Println(contact)
	}

	fmt.Println("Deleting all (truncating): ")
	err = contacts.DeleteAll()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
