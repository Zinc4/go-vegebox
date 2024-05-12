package admin

import (
	"errors"
	"math"
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"
)

type Usecase interface {
	GetUserPagination(page, pageSize int) ([]user.User, int, int, int, int, error)
	SearchUserByName(name string) ([]user.User, error)
	GetTransactionsPagination(page, pageSize int) ([]transaction.Transaction, int, int, int, int, error)
	SearchTransactionByName(name string) ([]transaction.Transaction, error)
	DeleteUserById(id int) (user.User, error)
	FindUserById(id int) (user.User, error)
	CreateProduct(input product.AddProductInput) (product.Product, error)
	UpdateProduct(inputID product.GetProductDetailInput, data product.AddProductInput) (product.Product, error)
	FindProductByID(id int) (product.Product, error)
	FindCategoryByID(id int) (product.Category, error)
	DeleteProductByID(id int) (product.Product, error)
	DeleteCategoryByID(id int) (product.Category, error)
	CreateCategory(input product.Category) (product.Category, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) GetUserPagination(page, pageSize int) ([]user.User, int, int, int, int, error) {
	totalUsers, err := u.repository.GetTotalUsers()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalUsers) / float64(pageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * pageSize
	users, err := u.repository.GetPaginatedUsers(offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}

	if page > 1 {
		prevPage = page - 1
	}

	return users, page, totalPages, nextPage, prevPage, nil

}

func (u *usecase) SearchUserByName(name string) ([]user.User, error) {
	users, err := u.repository.SearchUserByName(name)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *usecase) GetTransactionsPagination(page, pageSize int) ([]transaction.Transaction, int, int, int, int, error) {

	totalTransactions, err := u.repository.GetTotalTransaction()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalTransactions) / float64(pageSize)))

	if page < 1 {
		page = 1
	}

	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * pageSize
	transactions, err := u.repository.GetPaginatedTransaction(offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	}

	return transactions, page, totalPages, nextPage, prevPage, nil
}

func (u *usecase) SearchTransactionByName(name string) ([]transaction.Transaction, error) {

	userTransactions, err := u.repository.SearchTransactionByName(name)
	if err != nil {
		return userTransactions, err
	}
	return userTransactions, nil
}

func (u *usecase) DeleteUserById(id int) (user.User, error) {
	deletedUser, err := u.repository.GetUserByID(id)
	if err != nil {
		return user.User{}, err
	}

	return deletedUser, nil

}

func (u *usecase) FindUserById(id int) (user.User, error) {

	user, err := u.repository.FindUserById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, err

}

func (u *usecase) CreateProduct(input product.AddProductInput) (product.Product, error) {
	product := product.Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.Category.Name = input.Category.Name
	product.Category.ID = input.Category.ID

	newProduct, err := u.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil

}

func (u *usecase) UpdateProduct(inputID product.GetProductDetailInput, data product.AddProductInput) (product.Product, error) {
	product, err := u.repository.FindProductByID(inputID.ID)
	if err != nil {
		return product, err
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.Stock = data.Stock
	product.Category.ID = data.Category.ID
	product.Category.Name = data.Category.Name

	updatedProduct, err := u.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}

	return updatedProduct, nil

}

func (u *usecase) FindProductByID(id int) (product.Product, error) {

	product, err := u.repository.FindProductByID(id)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("product not found")
	}

	return product, err
}

func (u *usecase) DeleteProductByID(id int) (product.Product, error) {

	deletedProduct, err := u.repository.GetProductByID(id)
	if err != nil {
		return product.Product{}, err
	}

	return deletedProduct, nil
}

func (u *usecase) CreateCategory(input product.Category) (product.Category, error) {

	category := product.Category{}
	category.Name = input.Name

	newCategory, err := u.repository.SaveCategory(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil

}

func (u *usecase) FindCategoryByID(id int) (product.Category, error) {

	category, err := u.repository.FindByCategoryByID(id)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}

	return category, err
}

func (u *usecase) DeleteCategoryByID(id int) (product.Category, error) {

	deletedCategory, err := u.repository.GetCategoryByID(id)
	if err != nil {
		return product.Category{}, err
	}

	return deletedCategory, nil
}
