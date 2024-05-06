package product

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	CategoryID  int       `json:"category_id"`
	Price       int       `json:"price" gorm:"type:int"`
	Stock       int       `json:"stock" gorm:"type:int"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	DeleteAt    time.Time `json:"delete_at"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}

type Category struct {
	ID   int    `gorm:"primaryKey;" json:"id"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}
