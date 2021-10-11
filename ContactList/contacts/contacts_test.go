package contacts

import "testing"

const (
	hostname = "postgresql"
	hostPort = 5432
	user     = "postgres"
	password = "12345"
	dbName   = "postgres"
)

const (
	name    = "SomeName"
	newName = "NewName"
	phone   = 901234567
)

func TestCreateContact(t *testing.T) {
	err := InitDB(hostname, hostPort, user, password, dbName)
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()
	contact, err := NewContact(phone, name)
	if err != nil {
		t.Error(err)
	}
	if contact.getName() != name || contact.getPhone() != phone {
		t.Errorf("error: wrong details")
	}
}

func TestGetContact(t *testing.T) {
	err := InitDB(hostname, hostPort, user, password, dbName)
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()
	contact, err := NewContact(phone, name)
	if err != nil {
		t.Error(err)
	}
	if contact2, err := Get(contact.ID); err != nil {
		t.Error(err)
	} else if contact != contact2 {
		t.Errorf("expected %v, found %v", contact, contact2)
	}
}

func TestUpdateContact(t *testing.T) {
	err := InitDB(hostname, hostPort, user, password, dbName)
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()
	contact, err := NewContact(phone, name)
	if err != nil {
		t.Error(err)
	}
	contact2, err := UpdateContact(contact.ID, phone, newName)
	if err != nil {
		t.Error(err)
	}
	contact.Name = newName
	if contact != contact2 {
		t.Errorf("expected %v, found %v", contact, contact2)
	}
}

func TestDeleteContact(t *testing.T) {
	err := InitDB(hostname, hostPort, user, password, dbName)
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()
	contact, err := NewContact(phone, name)
	if err != nil {
		t.Error(err)
	}
	err = DeleteContact(contact.getId())
	if err != nil {
		t.Error(err)
	}
	contact, err = Get(contact.ID)
	if err == nil {
		t.Error("error expected")
	}
}

func TestGetAll(t *testing.T) {
	err := InitDB(hostname, hostPort, user, password, dbName)
	if err != nil {
		t.Error(err)
	}
	defer CloseDB()
	err = DeleteAll()
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		_, err = NewContact(phone, name)
		if err != nil {
			t.Error(err)
		}
	}
	result, err := GetAll()
	if err != nil {
		t.Error(err)
	}
	if n := len(result); n != 10 {
		t.Errorf("expected %v elements, got %v", 10, n)
	}
}
