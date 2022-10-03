package container

import (
	"swag-tutor/app/adapter/database"
	"swag-tutor/app/repository/person"
	"swag-tutor/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
}

func SetUpContainer() *Container {

	db := database.SetupDatabase()

	personRepo := person.NewPersonRepository(db)
	usecase := usecase.NewUsecase(personRepo)

	return &Container{
		Usecase: usecase,
	}
}