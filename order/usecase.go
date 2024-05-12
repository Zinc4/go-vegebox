package order

import (
	"mini-project/user"

	"github.com/google/uuid"
)

type Usecase interface {
	CreateOrder(order Order) (*Order, error)
	CreateOrderItems(orderItem []OrderItem) error
	GetOrderItemsByOrderID(orderID uuid.UUID) ([]OrderItem, error)
	GetOrderItemByID(uuid.UUID) (Order, error)
	GetOrdersByUserID(int) ([]Order, error)
	GetOrderSerializer(order Order) interface{}
	GetOrdersSerializerByUserID(user user.User) interface{}
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) CreateOrder(order Order) (*Order, error) {
	return u.repository.Create(order)
}

func (u *usecase) CreateOrderItems(orderItem []OrderItem) error {
	return u.repository.CreateOrderItems(orderItem)
}

func (u *usecase) GetOrderItemsByOrderID(orderID uuid.UUID) ([]OrderItem, error) {
	return u.repository.GetOrderItemsByOrderID(orderID)
}

func (u *usecase) GetOrdersByUserID(userID int) ([]Order, error) {
	return u.repository.GetOrdersByUserID(userID)
}

func (u *usecase) GetOrderItemByID(orderID uuid.UUID) (Order, error) {
	return u.repository.GetOrderItemByID(orderID)
}

func (u *usecase) GetOrderSerializer(order Order) interface{} {
	items, _ := u.GetOrderItemsByOrderID(order.ID)
	res_items := make([]ResponseOrderItem, len(items))
	for i, item := range items {
		res_items[i].Product = item.Product
		res_items[i].Quantity = item.Quantity
		res_items[i].Price = item.Product.Price * item.Quantity

	}

	res_order := ResponseOrder{
		ID:        order.ID,
		CreatedAt: order.CreatedAt,
		Status:    order.Status,
		Items:     res_items,
	}

	return res_order
}

func (u *usecase) GetOrdersSerializerByUserID(user user.User) interface{} {
	orders, _ := u.GetOrdersByUserID(user.ID)
	orders_ser := ResponseOrders{Orders: make([]ResponseOrder, len(orders))}
	for order_index, order := range orders {
		items, _ := u.GetOrderItemsByOrderID(order.ID)
		res_items := make([]ResponseOrderItem, len(items))
		for i, item := range items {
			res_items[i].Product = item.Product
			res_items[i].Quantity = item.Quantity
			res_items[i].Price = item.Product.Price * item.Quantity
		}
		order_ser := ResponseOrder{
			ID:        order.ID,
			CreatedAt: order.CreatedAt,
			Status:    order.Status,
			Items:     res_items,
		}
		orders_ser.Orders[order_index] = order_ser
	}
	return orders_ser
}
