package order

import (
	"mini-project/product"
	"time"

	"github.com/google/uuid"
)

type ResponseOrderItem struct {
	Product  product.Product `json:"product"`
	Quantity int             `json:"quantity"`
	Price    int             `json:"total_price"`
}

type ResponseOrder struct {
	ID        uuid.UUID           `json:"id"`
	Status    string              `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	Items     []ResponseOrderItem `json:"items"`
}

type ResponseOrders struct {
	Orders []ResponseOrder `json:"orders"`
}

type OrderFormatter struct {
	OrderId  uuid.UUID       `json:"order_id"`
	Status   string          `json:"status"`
	Name     string          `json:"name"`
	Amount   int             `json:"amount"`
	Product  product.Product `json:"product"`
	Quantity int             `json:"quantity"`
}

// type OrderProduct struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Price       int    `json:"price"`
// 	Stock       int    `json:"stock"`
// }

// func NewOrderProduct(product product.Product) OrderProduct {
// 	return OrderProduct{
// 		ID:          product.ID,
// 		Name:        product.Name,
// 		Description: product.Description,
// 		Price:       product.Price,
// 		Stock:       product.Stock,
// 	}
// }

func FormatOrder(order OrderItem) OrderFormatter {
	return OrderFormatter{
		OrderId:  order.Order.ID,
		Status:   order.Order.Status,
		Name:     order.Order.User.Name,
		Product:  order.Product,
		Amount:   order.Quantity * order.Product.Price,
		Quantity: order.Quantity,
	}
}

func FormatOrders(orders []OrderItem) []OrderFormatter {

	var orderFormatters []OrderFormatter

	for _, order := range orders {

		orderFormatters = append(orderFormatters, FormatOrder(order))
	}

	return orderFormatters
}
