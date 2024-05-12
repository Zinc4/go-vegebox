package order

import (
	"mini-project/product"
	"mini-project/user"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `gorm:"primaryKey;index" json:"id"`
	UserID    int       `json:"-" gorm:"index"`
	User      user.User `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	OrderID   uuid.UUID       `gorm:"primaryKey;index" json:"-"`
	Order     Order           `json:"order" gorm:"foreignKey:OrderID;references:ID"`
	ProductID int             `gorm:"primaryKey;index" json:"-"`
	Product   product.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int             `json:"quantity"`
	Price     int             `json:"price"`
}
