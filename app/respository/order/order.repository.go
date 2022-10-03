package order

import model "ass-david/app/models/mysql"

type OrderRepository interface {
	CreateOrder(order *model.Order) (result *model.Order, err error)
	CreateItem(item *model.Item) (result *model.Item, err error)
	FindAllOrder() (result *[]model.Order, err error)
	FindOneOrder(order_id int) (result *model.Order, err error)
	UpdateOrder(order_id int, order *model.Order) (result *model.Order, err error)
	UpdateItem(item_id int, item *model.Item) (result *model.Item, err error)
	DeleteOrder(order_id int) (err error)
}
