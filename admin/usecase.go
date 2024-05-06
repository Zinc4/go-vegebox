package admin

import (
	"errors"
	"math"
	"mini-project/product"
	"mini-project/user"
)

type Usecase interface {
	GetUserPagination(page, pageSize int) ([]user.User, int, int, int, int, error)
	SearchUserByName(name string) ([]user.User, error)
	DeleteUserById(id int) (user.User, error)
	FindUserById(id int) (user.User, error)
	CreateProduct(input product.AddProductInput) (product.Product, error)
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
	product.Category.ID = input.Category.ID
	product.Category.Name = input.Category.Name

	newProduct, err := u.repository.Save(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil

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
