package order

import (
	model "ass-david/app/models/mysql"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or orderRepository) CreateOrder(order *model.Order) (result *model.Order, err error) {
	err = or.db.Create(order).Error
	if err != nil {
		return result, err
	}
	return order, nil
}

func (or orderRepository) CreateItem(item *model.Item) (result *model.Item, err error) {
	err = or.db.Create(item).Error
	if err != nil {
		return result, err
	}
	return item, nil
}

func (or orderRepository) FindAllOrder() (result *[]model.Order, err error) {
	var order []model.Order
	err = or.db.Find(&order).Error
	for key, val := range order {
		var items []model.Item
		err = or.db.Model(&model.Item{}).Where("order_id = ?", val.OrderId).Find(&items).Error
		val.Items = items
		order[key] = val
	}
	result = &order
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}

func (or orderRepository) FindOneOrder(order_id int) (result *model.Order, err error) {
	var order model.Order
	var items []model.Item
	err = or.db.Where("order_id = ?", order_id).First(&order).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("order_id = ?", order_id).Find(&items).Error
	if err != nil {
		return result, err
	}
	order.Items = items
	fmt.Println(order, "order")
	result = &order
	return result, nil
}

func (or orderRepository) UpdateOrder(order_id int, order *model.Order) (result *model.Order, err error) {
	var orders model.Order
	err = or.db.Model(&orders).Clauses(clause.Returning{}).Where("order_id = ?", order_id).Updates(order).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("order_id = ?", order_id).First(&orders).Error
	if err != nil {
		return result, err
	}
	fmt.Println(orders, "orders")
	return &orders, nil
}

func (or orderRepository) UpdateItem(item_id int, item *model.Item) (result *model.Item, err error) {
	var items model.Item
	err = or.db.Model(&items).Clauses(clause.Returning{}).Where("item_id = ?", item_id).Updates(item).Error
	if err != nil {
		return result, err
	}
	err = or.db.Where("item_id = ?", item_id).First(&items).Error
	if err != nil {
		return result, err
	}
	fmt.Println(items, "items")
	return &items, nil
}

func (or orderRepository) DeleteOrder(order_id int) (err error) {
	err = or.db.Where("order_id = ?", order_id).Delete(&model.Item{}).Error
	if err != nil {
		return err
	}
	err = or.db.Delete(&model.Order{}, order_id).Error
	if err != nil {
		return err
	}
	return nil
}
