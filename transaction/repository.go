package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetByUserId(id int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	GetByCode(code string) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByUserId(id int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("User").Preload("Order").Where("user_id = ?", id).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByCode(code string) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("code = ?", code).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
