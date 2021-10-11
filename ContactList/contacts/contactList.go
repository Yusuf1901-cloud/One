package contacts

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const max = 1000

type Contact struct {
	ID, PhoneNumber int
	Name            string
}

func (contact Contact) getId() int {
	return contact.ID
}

func (contact Contact) getPhone() int {
	return contact.PhoneNumber
}

func (contact Contact) getName() string {
	return contact.Name
}

func (contact Contact) String() string {
	return fmt.Sprintf("%v-- Name: %v and PhoneNumber is %v", contact.ID, contact.Name, contact.PhoneNumber)
}

var database *sql.DB

func InitDB(hostname string, hostPort int, user, password, dbName string) (err error) {
	err = nil
	connectingstring := fmt.Sprintf("port = %d host = %s user = %s"+
		"password = %s dbname = %s sslmode = disable",
		hostPort, hostname, user, password, dbName)
	database, err = sql.Open("postgres", connectingstring)
	if err != nil {
		return
	}
	_, err = database.Exec("CREATE TABLE IF NOT EXIST contacts(" +
		"id INT UNIQUE NOT NULL PRIMARY KEY, phone INT NOT NULL, name TEXT NOT NULL);")
	return
}

func CloseDB(err error) {
	err = database.Close()
	return
}

func NewContact(phoneNumber int, name string) (Contact, error) {
	var id int64
	err := database.QueryRow("INSERT INTO contact(name, phonenumber) VALUES($1, $2) id;",
		name, phoneNumber).Scan(&id)
	if err != nil {
		return Contact{}, err
	}
	contact := Contact{int(id), phoneNumber, name}
	return contact, err
}
func Get(id int) (Contact, error) {
	result, err := database.Query("SELECT *FROM contact WHERE id = $1;", id)

	if err != nil {
		return Contact{}, err
	}
	contact := Contact{}
	if result.Next() {
		err = result.Scan(&contact.ID, &contact.PhoneNumber, &contact.Name)
		return contact, err
	}
	return contact, errors.New("contact not found")
}

func UpdateContact(id, newPhone int, newName string) (Contact, error) {
	_, err := Get(id)

	if err != nil {
		return Contact{}, err
	}
	_, err = database.Exec("UPDATE contact SET phonenumber = $1, name = $2 WHERE id =$3;", newPhone, newName, id)
	contact := Contact{id, newPhone, newName}
	return contact, err
}
func DeleteContact(id int) error {
	_, err := Get(id)
	if err != nil {
		return err
	}
	_, err = database.Exec("DELETE FROM contact WHERE id = $1;", id)

	return err
}

func GetAll() ([]Contact, error) {
	result := make([]Contact, 0, max)
	rows, err := database.Query("SELECT *FROM contact;")

	if err != nil {
		return result, err
	}
	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.PhoneNumber, &contact.Name)
		if err != nil {
			return result, err
		}
		result = append(result, contact)
	}

	return result, err

}

func DeleteAll() error {
	_, err := database.Exec("TRUNCATE TABLE contact;")
	return err
}
