package request

type GetOnePersonRequestParam struct {
	Id int `uri:"id" binding:"required"`
}

type CreatePersonRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name,omitempty"`
}

type UpdatePersonRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
}

type UpdatePersonRequestParam struct {
	Id int `uri:"id" binding:"required"`
}

type DeletePersonRequestParam struct {
	Id int `uri:"id" binding:"required"`
}