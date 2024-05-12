package transaction

import (
	"mini-project/order"
	"mini-project/user"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         int         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int         `json:"user_id" gorm:"index"`
	OrderID    uuid.UUID   `json:"order_id" gorm:"index"`
	Status     string      `json:"status"`
	Code       string      `json:"code"`
	Amount     int         `json:"amount"`
	PaymentURL string      `json:"payment_url" gorm:"type:varchar(255)"`
	User       user.User   `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Order      order.Order `json:"order" gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  time.Time   `json:"deleted_at"`
}
