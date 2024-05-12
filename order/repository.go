package order

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(order Order) (*Order, error)
	Update(order Order) (Order, error)
	GetOrderByID(uuid.UUID) (Order, error)
	CreateOrderItems(orderItem []OrderItem) error
	GetOrderItemsByOrderID(orderID uuid.UUID) ([]OrderItem, error)
	GetOrdersByUserID(int) ([]Order, error)
	GetOrderItemByID(uuid.UUID) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(order Order) (*Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return &order, err
	}
	return &order, nil
}

func (r *repository) Update(order Order) (Order, error) {

	result := r.db.Where("Id = ?", order.ID).Updates(&order)
	if result.Error != nil {
		return order, result.Error
	}
	if result.RowsAffected == 0 {
		return order, errors.New("order not found")
	}
	return order, nil
}

func (r *repository) GetOrderByID(orderID uuid.UUID) (Order, error) {

	var order Order
	err := r.db.Preload("User").First(&order, orderID).Error
	if err != nil {
		return order, err
	}

	return order, nil

}

func (r *repository) CreateOrderItems(orderItems []OrderItem) error {

	return r.db.Create(&orderItems).Error
}

func (r *repository) GetOrderItemsByOrderID(orderID uuid.UUID) ([]OrderItem, error) {
	var orderItems []OrderItem
	err := r.db.Preload("Product").Preload("Product.Category").Preload("Order").Preload("Order.User").Where("order_id = ?", orderID).Find(&orderItems).Error
	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func (r *repository) GetOrdersByUserID(userID int) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *repository) GetOrderItemByID(orderID uuid.UUID) (Order, error) {

	var order Order
	err := r.db.Preload("User").Where("id = ?", orderID).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil

}
