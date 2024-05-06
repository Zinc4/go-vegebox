package admin

import (
	"mini-project/product"
	"mini-project/user"

	"gorm.io/gorm"
)

type Repository interface {
	GetTotalUsers() (int64, error)
	GetPaginatedUsers(offset int, limit int) ([]user.User, error)
	SearchUserByName(name string) ([]user.User, error)
	GetUserByID(userID int) (user.User, error)
	FindUserById(userID int) (user.User, error)
	Save(product product.Product) (product.Product, error)
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

func (r *repository) Save(product product.Product) (product.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) SaveCategory(category product.Category) (product.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
