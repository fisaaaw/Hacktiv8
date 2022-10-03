package container

import (
	"ass-david/app/adapter/database"
	order "ass-david/app/respository/order"
	"ass-david/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
}

func SetupContainer() Container {
	db := database.SetupDatabase().DB

	repo := order.NewOrderRepository(db)
	usecase := usecase.NewUsecase(repo)

	return Container{
		Usecase: usecase,
	}

}
