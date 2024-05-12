package admin

import (
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"

	"gorm.io/gorm"
)

type Repository interface {
	GetTotalUsers() (int64, error)
	GetPaginatedUsers(offset int, limit int) ([]user.User, error)
	SearchUserByName(name string) ([]user.User, error)
	GetUserByID(userID int) (user.User, error)
	FindUserById(userID int) (user.User, error)
	GetTotalTransaction() (int64, error)
	GetPaginatedTransaction(offset int, limit int) ([]transaction.Transaction, error)
	SearchTransactionByName(name string) ([]transaction.Transaction, error)
	GetTotalTransactionByName(name string) (int64, error)
	Save(product product.Product) (product.Product, error)
	Update(product product.Product) (product.Product, error)
	FindProductByID(ID int) (product.Product, error)
	FindByCategoryByID(ID int) (product.Category, error)
	GetProductByID(ID int) (product.Product, error)
	GetCategoryByID(ID int) (product.Category, error)
	SaveCategory(category product.Category) (product.Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTotalUsers() (int64, error) {
	var total int64
	err := r.db.Model(&user.User{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetPaginatedUsers(offset int, limit int) ([]user.User, error) {
	var users []user.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) SearchUserByName(name string) ([]user.User, error) {
	var users []user.User
	err := r.db.Where("name = ?", name).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) GetUserByID(userID int) (user.User, error) {
	var user user.User
	err := r.db.Where("id = ?", userID).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindUserById(userID int) (user.User, error) {
	var user user.User
	err := r.db.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetTotalTransaction() (int64, error) {
	var total int64
	err := r.db.Model(&transaction.Transaction{}).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetPaginatedTransaction(offset int, limit int) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	err := r.db.Offset(offset).Limit(limit).Preload("User").Preload("Order").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) SearchTransactionByName(name string) ([]transaction.Transaction, error) {
	var userTransactions []transaction.Transaction
	err := r.db.Preload("User").Preload("Order").Joins("JOIN users ON transactions.user_id = users.id").Where("users.name LIKE ?", "%"+name+"%").Find(&userTransactions).Error
	if err != nil {
		return userTransactions, err
	}
	return userTransactions, nil
}

func (r *repository) GetTotalTransactionByName(name string) (int64, error) {
	var total int64
	var userTransactions []transaction.Transaction
	err := r.db.Preload("User").Preload("Order").Joins("JOIN users ON transactions.user_id = users.id").Where("users.name LIKE ?", "%"+name+"%").Count(&total).Find(&userTransactions).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) Save(product product.Product) (product.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(product product.Product) (product.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindProductByID(ID int) (product.Product, error) {
	var product product.Product
	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByCategoryByID(ID int) (product.Category, error) {
	var category product.Category
	err := r.db.Where("id = ?", ID).Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) GetProductByID(ID int) (product.Product, error) {
	var product product.Product
	err := r.db.Where("id = ?", ID).Delete(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) GetCategoryByID(ID int) (product.Category, error) {
	var category product.Category
	err := r.db.Where("id = ?", ID).Delete(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) SaveCategory(category product.Category) (product.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
