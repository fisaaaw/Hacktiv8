package transport

import (
	// "strconv"
	"swag-tutor/app/usecase"
	"swag-tutor/app/usecase/request"

	"github.com/gin-gonic/gin"
)

type tp struct {
	usecase usecase.Usecase
}

func NewTransport(usecase usecase.Usecase) *tp {
	return &tp{
		usecase: usecase,
	}
}

// Get All Persons
// @Summary Mendapatkan semua orang yang telah terdaftar
// @Description Mendapatkan semua orang yang telah terdaftar
// @Tags Person
// @Accept  */*
// @Produce  json
// @Success 200 {object} response.GetAllPersonsResponse
// @Failure 500 {object} string "error"
// @Router /v1/person [get]
func (transprt tp) GetAllPersons(c *gin.Context) {
	res, err := transprt.usecase.GetAllPersons()
	if err != nil {
		c.JSON(res.HttpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(res.HttpCode, res.Payload)
}

// Get One Person
// @Summary Mendapatkan satu orang yang telah terdaftar dengan id
// @Description Mendapatkan satu orang yang telah terdaftar dengan id
// @Tags Person
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} response.GetOnePersonResponse
// @Failure 500 {object} string "error"
// @Router /v1/person/{id} [get]
func (transport tp) GetOnePerson(c *gin.Context) {
	var person request.GetOnePersonRequestParam
	err := c.ShouldBindUri(&person)
	// id := c.Param("id")
	// aid, err := strconv.Atoi(id)
	// person.Id = aid
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something error on bind uri",
		})
		return
	}

	res, err := transport.usecase.GetOnePerson(person.Id)
	if err != nil {
		c.JSON(res.HttpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(res.HttpCode, res.Payload)
}

// Create One Person
// @Summary Mendaftarkan/membuat satu orang
// @Description Mendaftarkan/membuat satu orang
// @Tags Person
// @Accept  */*
// @Produce  json
// @Param data body request.CreatePersonRequest true "Person"
// @Success 200 {object} response.CreatePersonResponse
// @Failure 500 {object} string "error"
// @Router /v1/person [post]
func (transport tp) CreatePerson(c *gin.Context) {
	var person request.CreatePersonRequest
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something field is missing",
		})
		return
	}

	res, err := transport.usecase.CreatePerson(&person)
	if err != nil {
		c.JSON(res.HttpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(res.HttpCode, res.Payload)
}

// Update One Person
// @Summary Mengupdate satu orang yang sudah terbuat
// @Description Mengupdate satu orang yang sudah terbuat
// @Tags Person
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Param data body request.UpdatePersonRequest true "Person"
// @Success 200 {object} response.UpdatePersonResponse
// @Failure 500 {object} string "error"
// @Router /v1/person/{id} [put]
func (transport tp) UpdatePersonRequest(c *gin.Context) {
	var person request.UpdatePersonRequest
	var identify request.UpdatePersonRequestParam
	err := c.ShouldBindUri(&identify)
	// id := c.Param("id")
	// aid, err := strconv.Atoi(id)
	// identify.Id = aid
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something wrong on bind id",
		})
		return
	}

	err = c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something wrong on bind body",
		})
		return
	}
	res, err := transport.usecase.UpdatePerson(identify.Id, &person)
	if err != nil {
		c.JSON(res.HttpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(res.HttpCode, res.Payload)
}

// Delete
// @Summary Menghapus user yang sudah dibuat berdasarkan id
// @Description Menghapus user yang sudah dibuat berdasarkan id
// @Tags Person
// @Accept  */*
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} response.DeletePersonResponse
// @Failure 400 {object} string "error"
// @Router /v1/person/{id} [delete]
func (transport tp) DeletePerson(c *gin.Context) {
	var identify request.DeletePersonRequestParam
	err := c.ShouldBindUri(&identify)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something wrong on bind id",
		})
		return
	}

	res, err := transport.usecase.DeletePerson(identify.Id)
	if err != nil {
		c.JSON(res.HttpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(res.HttpCode, res.Payload)
}
