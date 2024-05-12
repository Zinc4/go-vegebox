package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	FindAllCategory() ([]Category, error)
	FindByCategory(id int) ([]Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Category").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	err := r.db.Preload("Category").Where("id = ?", ID).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindAllCategory() ([]Category, error) {
	var categories []Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *repository) FindByCategory(id int) ([]Product, error) {
	var products []Product
	err := r.db.Where("category_id = ?", id).Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}
