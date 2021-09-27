package contacts

import "testing"

const name = "Dadaxon"
const id = 31019
const phone = 998541756
const email = "farboborahimov@gmail.com"

const n = 5

func TestNewContact(t *testing.T) {
	contact := NewContact(id, phone, name, email)

	if contact.Name != name {
		t.Errorf("\nExpected %s, found: %s\n", name, contact.Name)
	}
	if contact.Phone != phone {
		t.Errorf("\nExpected: %v, found: %v\n", phone, contact.Phone)
	}
}
func TestNewContactList(t *testing.T) {
	//contact := NewContact(id, phone, name, email)
	contactList := NewContactList(n)

	if topil := contactList.size; topil != n {
		t.Error("\nNumber of Objects is not correct!!\n")
	}
}
func TestCreateContactList(t *testing.T) {
	contactList := NewContactList(n)
	contact := NewContact(id, phone, name, email)

	prob := contactList.Create(contact)
	if prob != nil {
		t.Errorf("\nCreating contact : %v\n", prob)
	}

}

func TestGetContact(t *testing.T) {
	contactList := NewContactList(n)
	contact := NewContact(id, phone, name, email)

	for i := 0; i < n; i++ {
		contactList.Create(contact)
	}

	if err := contactList.get(0); err != contact {
		t.Errorf("\nError contact found: %v, expected : %v\n", err, contact)
	}
}

func TestUpdateContact(t *testing.T) {
	contactList := NewContactList(n)
	contact := NewContact(id, phone, name, email)

	for i := 0; i < n; i++ {
		contactList.Create(contact)
	}

	newContact := NewContact(31219, 331254851, "Farrux", "farruxsal@gmail.com")

	if err := contactList.update(0, newContact); err == nil {
		t.Errorf("Error : %v\n", err)
	}
	if found := contactList.get(0); found != newContact {
		t.Errorf("Expected: %v, found: %v\n", newContact, found)
	}
}

func TestGetAll(t *testing.T) {
	contact := NewContact(id, phone, name, email)
	contactList := NewContactList(n)

	for i := 0; i < n; i++ {
		contactList.Create(contact)
	}

	if len(contactList.getAll()) != n {
		t.Errorf("Expected length : %v, but now: %v\n", n, len(contactList.getAll()))
	}
}
func TestDeleteContact(t *testing.T) {
	contactList := NewContactList(n)
	contact := NewContact(id, phone, name, email)

	for i := 0; i < n; i++ {
		contactList.Create(contact)
	}

	if err := contactList.delete(0); err != nil {
		t.Errorf("Error, value: %v\n", err)
	}

	if contactList.n != n-1 {
		t.Errorf("Expected %v, found: %v\n", n-1, contactList.n)
	}
}
