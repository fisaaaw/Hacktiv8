package usecase

import (
	"ass-david/app/usecase/request"
	"ass-david/app/usecase/response"
)

type Usecase interface {
	CreateOrder(body *request.CreateOrderRequest) (result *response.CreateOrderResponse, http_code int, err error)
	FindAllOrder() (result *response.FindAllOrderResponse, http_code int, err error)
	FindOneOrder(id int) (result *response.FindOneOrderResponse, http_code int, err error)
	UpdateOrder(id int, body *request.UpdateOrderRequest) (result *response.UpdateOrderResponse, http_code int, err error)
	DeleteOrder(id int) (result *response.DeleteOrderResponse, http_code int, err error)
}
