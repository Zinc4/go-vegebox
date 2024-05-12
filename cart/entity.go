package cart

import (
	"mini-project/product"
	"time"
)

type Cart struct {
	ID        int       `gorm:"index;primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CartItem struct {
	CartID    int             `gorm:"index;primaryKey" json:"-"`
	Cart      Cart            `gorm:"foreignKey:CartID;references:ID" json:"cart"`
	ProductID int             `gorm:"index;primaryKey" json:"-"`
	Product   product.Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
}
