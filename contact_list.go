package main

import "fmt"

type Contact struct {
	ID, Phone                            int
	FirstName, LastName, Email, Position string
}

type ContactList struct {
	contacts []Contact
	n        int
}

func (new *Contact) NewContact(a, b int, c, d, e, f string) {
	new.ID = a
	new.Phone = b
	new.FirstName = c
	new.LastName = d
	new.Email = e
	new.Position = f
}

func (contactList *ContactList) Create(contact *Contact) {

	contactList.contacts = append(contactList.contacts, *contact)

}

func get(con *Contact) string {

	return fmt.Sprintf(
		"%s %s's ID: %v and Email: %s. Now you are work at %s position. Tel: %v \n",
		con.FirstName,
		con.LastName,
		con.ID,
		con.Email,
		con.Position,
		con.Phone,
	)

}

func (contactlist *ContactList) update(index int, newContact Contact) {
	contactlist.contacts[index] = newContact
}

func (contactList ContactList) getAll() string {
	que := ""
	for i, con := range contactList.contacts {
		que += fmt.Sprintf("\n%v: %v\n\n", i, con)
	}
	return que
}

func (contactList *ContactList) delete(index int) {
	if index < len(contactList.contacts) {
		contactList.contacts = append(contactList.contacts[:index],
			contactList.contacts[index+1:]...)
	}
}

func main() {
	k := Contact{}
	l := ContactList{}

	k.NewContact(31219, 985471230, "Gafur", "Tursunov", "gtursunov@gmail.com", "Teacher")
	l.Create(&k)

	fmt.Println(get(&k))

	k.NewContact(31519, 987876230, "Gafur", "Tursunov", "gtursunov@gmail.com", "Teacher")
	l.Create(&k)

	fmt.Println(l.getAll())

	fmt.Println("\nAfter updated:")

	k.NewContact(31319, 902100410, "Amon", "Turdiyev", "jxalilov@gmail.com", "Cheif")
	l.update(1, k)
	fmt.Println(l.getAll())

	var n int
	fmt.Scanln(&n)

	fmt.Printf("\n Now deleting %v index of Contactlist\n", n)
	l.delete(n)

	fmt.Println("\nAfter deleting:")
	fmt.Println(l.getAll())
}
