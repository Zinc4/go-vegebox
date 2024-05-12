package cart

import (
	"mini-project/product"
	"time"
)

type CartItemFormatter struct {
	Product   product.Product `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
}

func FormatProduct(item CartItem) CartItemFormatter {

	return CartItemFormatter{

		Product:   item.Product,
		Quantity:  item.Quantity,
		CreatedAt: item.CreatedAt,
	}
}

type CartsFormatter struct {
	ID        int `json:"id"`
	Items     []CartItemFormatter
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
