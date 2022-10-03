package person

import model "swag-tutor/app/models/sql"

type PersonRepository interface {
	GetAllPersons() (result *[]model.Person, err error)
	GetOnePerson(id int) (result *model.Person, err error)
	CreatePerson(person *model.Person) (result *model.Person, err error)
	UpdatePerson(person *model.Person, id int) (result string, err error)
	DeletePerson(id int) (result string, err error)
}
