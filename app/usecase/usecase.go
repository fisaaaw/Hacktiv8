package usecase

import (
	"swag-tutor/app/usecase/request"
	"swag-tutor/app/usecase/response"
)

type Usecase interface {
	GetAllPersons() (response *response.GetAllPersonsResponse, err error)
	GetOnePerson(id int) (response *response.GetOnePersonResponse, err error)
	CreatePerson(body *request.CreatePersonRequest) (response *response.CreatePersonResponse, err error)
	UpdatePerson(id int, body *request.UpdatePersonRequest) (response *response.UpdatePersonResponse, err error)
	DeletePerson(id int) (response *response.DeletePersonResponse, err error)
}
