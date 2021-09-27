package contacts

import (
	"errors"
)

type Contact struct {
	ID, Phone   int
	Name, Email string
}

type ContactList struct {
	contacts []Contact
	n        int
	size     int
}

func NewContact(a, b int, c, d string) Contact {
	return Contact{a, b, c, d}

}

func NewContactList(size int) ContactList {
	return ContactList{make([]Contact, 0, size), 0, size}
}

func (contactList *ContactList) Create(contact Contact) error {
	if contactList.n < contactList.size {
		contactList.contacts = append(contactList.contacts, contact)
		contactList.n++
		return nil
	}
	return errors.New("rejected! range out off.\n")

}

func (contactList *ContactList) get(index int) Contact {
	if index < contactList.n {
		return contactList.contacts[index]
	}
	return Contact{}
}

func (contactlist *ContactList) update(index int, newContact Contact) error {
	if index < contactlist.n {
		contactlist.contacts[index] = newContact
		return nil
	}
	return errors.New("the contact isn't found with given index!\n")
}

func (contactList ContactList) getAll() []Contact {
	return contactList.contacts
}

func (contactList *ContactList) delete(index int) error {
	if index < len(contactList.contacts) {
		contactList.contacts = append(contactList.contacts[:index],
			contactList.contacts[index+1:]...)
		return nil
	}
	return errors.New("error: index out of range\n")
}

/*
func main() {
	k := Contact{}
	l := ContactList{}

	NewContact(31219, 985471230, "Gafur", "gtursunov@gmail.com")
	l.Create(&k)

	fmt.Println(get(&k))

	NewContact(31519, 987876230, "Gafur", "gtursunov@gmail.com")
	l.Create(&k)

	fmt.Println(l.getAll())

	fmt.Println("\nAfter updated:")

	NewContact(31319, 902100410, "Amon","jxalilov@gmail.com")
	l.update(1, k)
	fmt.Println(l.getAll())

	var n int
	fmt.Scanln(&n)

	fmt.Printf("\n Now deleting %v index of Contactlist\n", n)
	l.delete(n)

	fmt.Println("\nAfter deleting:")
	fmt.Println(l.getAll())
}
*/
