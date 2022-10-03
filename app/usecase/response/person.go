package response

import model "swag-tutor/app/models/sql"

type GetAllPersonsResponse struct {
	
	Payload *[]model.Person `json:"payload"`
	HttpCode int `json:"http_code"`
}

type GetOnePersonResponse struct {
	Payload *model.Person `json:"payload"`
	HttpCode int `json:"http_code"`
}

type CreatePersonResponse struct {
	Payload model.Person `json:"payload"`
	HttpCode int `json:"http_code"`
}

type UpdatePersonResponse struct {
	Payload string `json:"payload"`
	HttpCode int `json:"http_code"`
}

type DeletePersonResponse struct {
	Payload string `json:"payload"`
	HttpCode int `json:"http_code"`
}