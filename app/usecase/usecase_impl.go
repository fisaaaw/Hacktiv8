package usecase

import (
	"fmt"
	model "swag-tutor/app/models/sql"
	"swag-tutor/app/repository/person"
	"swag-tutor/app/usecase/request"
	"swag-tutor/app/usecase/response"
)

type usecase struct {
	person person.PersonRepository
}

func NewUsecase(person person.PersonRepository) Usecase {
	return &usecase{
		person: person,
	}
}

func (uc usecase) GetAllPersons() (resp *response.GetAllPersonsResponse, err error) {
	var result response.GetAllPersonsResponse
	res, err := uc.person.GetAllPersons()
	if err != nil {
		result.HttpCode = 500
		resp = &result
		return resp, err
	}
	result.HttpCode = 200
	result.Payload = res
	resp = &result
	return resp, nil
}

func (uc usecase) GetOnePerson(id int) (resp *response.GetOnePersonResponse, err error) {
	var result response.GetOnePersonResponse
	if id < 1 {
		result.HttpCode = 500
		resp = &result
		return resp, fmt.Errorf("%s", "ID value is empty")
	}
	res, err := uc.person.GetOnePerson(id)
	if err != nil {
		result.HttpCode = 500
		resp = &result
		return resp, err
	}
	result.HttpCode = 200
	result.Payload = res
	resp = &result
	return resp, nil
}

func (uc usecase) CreatePerson(body *request.CreatePersonRequest) (resp *response.CreatePersonResponse, err error) {
	var result response.CreatePersonResponse
	createPerson := &model.Person{
		FirstName: body.FirstName,
		LastName: body.LastName,
	}
	fmt.Println(createPerson)
	res, err := uc.person.CreatePerson(createPerson)
	if err != nil {
		result.HttpCode = 500
		resp = &result
		return resp, err
	}
	result.HttpCode = 200
	result.Payload = *res
	resp = &result
	return resp, nil
}

func (uc usecase) UpdatePerson(id int, body *request.UpdatePersonRequest) (resp *response.UpdatePersonResponse, err error) {
	var result response.UpdatePersonResponse
	updatePerson := &model.Person{
		FirstName: body.FirstName,
		LastName: body.LastName,
	}
	fmt.Println(id, "id")
	fmt.Println(updatePerson, "body")
	res, err := uc.person.UpdatePerson(updatePerson, id)
	if err != nil {
		result.HttpCode = 500
		resp = &result
		return resp, err
	}
	result.HttpCode = 200
	result.Payload = res
	resp = &result
	return resp, nil
}

func (uc usecase) DeletePerson(id int) (resp *response.DeletePersonResponse, err error) {
	var result response.DeletePersonResponse
	res, err := uc.person.DeletePerson(id)
	if err != nil {
		result.HttpCode = 500
		resp = &result
		return resp, err
	}
	result.HttpCode = 200
	result.Payload = res
	resp = &result
	return resp, nil
}