package admin

import (
	"errors"
	"mini-project/admin/mocks"
	"mini-project/product"
	"mini-project/transaction"
	"mini-project/user"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUserPagination(t *testing.T) {
	repository := mocks.NewRepository(t)
	usecase := NewUsecase(repository)

	var userList = []user.User{
		{
			ID:         1,
			Name:       "Rian",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
		{
			ID:         2,
			Name:       "Ihsan",
			Email:      "ihsanganteng@gmail.com",
			Password:   "ihsan12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
	}

	t.Run("Success get user pagination", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get total user failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(0), errors.New("get total users error")).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get total users error")
		assert.Nil(t, users)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Get paginated users failed", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return([]user.User{}, errors.New("get paginated users error")).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get paginated users error")
		assert.Nil(t, users)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 5
		var pageSize = 10

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 0, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1

		repository.On("GetTotalUsers").Return(int64(2), nil).Once()
		repository.On("GetPaginatedUsers", 1, pageSize).Return(userList, nil).Once()

		users, totalPages, currentPage, nextPage, prevPage, err := usecase.GetUserPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, userList, users)
		assert.Equal(t, userList[0].Name, users[0].Name)
		assert.Equal(t, userList[1].Name, users[1].Name)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)
	})

}

func TestSearchUserByName(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var userList = []user.User{
		{
			ID:         1,
			Name:       "Rian",
			Email:      "rianganteng@gmail.com",
			Password:   "rian12345",
			Avatar:     "www.cloudinary.com/avatar",
			Role:       "user",
			IsVerified: false,
		},
	}

	t.Run("Success get user", func(t *testing.T) {
		repository.On("SearchUserByName", "Rian").Return(userList, nil).Once()

		result, err := usecase.SearchUserByName("Rian")
		assert.Nil(t, err)
		assert.Equal(t, userList, result)
		repository.AssertExpectations(t)
	})

	t.Run("get user failed", func(t *testing.T) {
		repository.On("SearchUserByName", "Rian").Return([]user.User{}, errors.New("get user error")).Once()

		result, err := usecase.SearchUserByName("Rian")
		assert.Error(t, err)
		assert.Equal(t, []user.User{}, result)

		repository.AssertExpectations(t)
	})
}

func TestGetTransactionsPagination(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var mockTransaction = []transaction.Transaction{
		{
			ID:         1,
			UserID:     3,
			OrderID:    uuid.Nil,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         2,
			UserID:     3,
			OrderID:    uuid.Nil,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("Succes get all transaction users", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalTransaction").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransaction", 0, pageSize).Return(mockTransaction, nil).Once()
		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("Failed Get total transaction user ", func(t *testing.T) {
		var page = 1
		var pageSize = 10

		repository.On("GetTotalTransaction").Return(int64(0), errors.New("get total users transactions error")).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Error(t, err)
		assert.EqualError(t, err, "get total users transactions error")
		assert.Nil(t, transactions)
		assert.Empty(t, totalPages)
		assert.Empty(t, currentPage)
		assert.Empty(t, nextPage)
		assert.Empty(t, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page = 0", func(t *testing.T) {
		var page = 0
		var pageSize = 10

		repository.On("GetTotalTransaction").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransaction", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 5
		var pageSize = 10

		repository.On("GetTotalTransaction").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransaction", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 1, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > totalPage", func(t *testing.T) {
		var page = 1
		var pageSize = 1

		repository.On("GetTotalTransaction").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransaction", 0, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 1, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 2, nextPage)
		assert.Equal(t, 0, prevPage)
		repository.AssertExpectations(t)
	})

	t.Run("page > 1", func(t *testing.T) {
		var page = 2
		var pageSize = 1

		repository.On("GetTotalTransaction").Return(int64(2), nil).Once()
		repository.On("GetPaginatedTransaction", 1, pageSize).Return(mockTransaction, nil).Once()

		transactions, totalPages, currentPage, nextPage, prevPage, err := usecase.GetTransactionsPagination(page, pageSize)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction, transactions)
		assert.Equal(t, mockTransaction[0].ID, transactions[0].ID)
		assert.Equal(t, mockTransaction[1].ID, transactions[1].ID)
		assert.Equal(t, 2, totalPages)
		assert.Equal(t, 2, currentPage)
		assert.Equal(t, 0, nextPage)
		assert.Equal(t, 1, prevPage)
		repository.AssertExpectations(t)
	})
}

func TestSearchTransactionByName(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var transactionList = []transaction.Transaction{
		{
			ID:         1,
			UserID:     3,
			OrderID:    uuid.Nil,
			Amount:     10000,
			Status:     "paid",
			Code:       "48011",
			PaymentURL: "https://app.sandbox.midtrans.com/snap/v3/redirection/7a45765e-a0f3-4bfb-9b3f-a58560f09611",
			User:       user.User{Name: "user"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	t.Run("Success get transaction list by username", func(*testing.T) {
		repository.On("SearchTransactionByName", "user").Return(transactionList, nil).Once()

		result, err := usecase.SearchTransactionByName("user")
		assert.Nil(t, err)
		assert.Equal(t, transactionList, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed get transaction list by username", func(*testing.T) {
		repository.On("SearchTransactionByName", "user").Return([]transaction.Transaction{}, errors.New("Failed Get transaction user by user name")).Once()

		result, err := usecase.SearchTransactionByName("user")
		assert.NotNil(t, err)
		assert.Equal(t, []transaction.Transaction{}, result)

		repository.AssertExpectations(t)
	})
}

func TestDeleteUserById(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var user = user.User{
		ID:         1,
		Name:       "Rian",
		Email:      "rianganteng@gmail.com",
		Password:   "rian12345",
		Avatar:     "www.cloudinary.com/avatar",
		Role:       "user",
		IsVerified: true,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	t.Run("success delete user", func(t *testing.T) {
		repository.On("GetUserByID", user.ID).Return(user, nil).Once()
		result, err := usecase.DeleteUserById(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, result.ID)
		repository.AssertExpectations(t)

	})

	t.Run("Failed to delete user", func(t *testing.T) {
		repository.On("GetUserByID", user.ID).Return(user, errors.New("User not found")).Once()

		_, err := usecase.DeleteUserById(user.ID)
		assert.Error(t, err)
		assert.Equal(t, "User not found", err.Error())
		repository.AssertExpectations(t)

	})
}

func TestFindUserById(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	mockUser := user.User{
		ID:         1,
		Name:       "Rian",
		Email:      "rianganteng@gmail.com",
		Password:   "rian12345",
		Avatar:     "www.cloudinary.com/avatar",
		Role:       "user",
		IsVerified: true,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}

	t.Run("Success get user id", func(t *testing.T) {
		repository.On("FindUserById", mockUser.ID).Return(mockUser, nil).Once()
		result, err := usecase.FindUserById(mockUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.ID, result.ID)
		repository.AssertExpectations(t)

	})

	t.Run("Failed get user id", func(t *testing.T) {
		repository.On("FindUserById", mockUser.ID).Return(mockUser, errors.New("no user found with that ID")).Once()
		_, err := usecase.FindUserById(mockUser.ID)
		assert.Error(t, err)
		assert.Equal(t, 1, mockUser.ID, "no user found with that ID")
		repository.AssertExpectations(t)
	})
}

func TestCreateProduct(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	categoryData := product.Category{
		ID:   1,
		Name: "test",
	}

	var mockProduct = product.Product{
		Name:        "test",
		Description: "test",
		Price:       10000,
		Stock:       10,
		Category:    categoryData,
	}

	var input = product.AddProductInput{
		Name:        "test",
		Description: "test",
		Price:       10000,
		Stock:       10,
		Category:    categoryData,
	}

	t.Run("Success create product", func(t *testing.T) {
		repository.On("Save", mockProduct).Return(mockProduct, nil).Once()
		result, err := usecase.CreateProduct(input)
		assert.NoError(t, err)
		assert.Equal(t, mockProduct, result)
		repository.AssertExpectations(t)

	})

	t.Run("Failed create product", func(t *testing.T) {
		repository.On("Save", mockProduct).Return(mockProduct, errors.New("error")).Once()
		_, err := usecase.CreateProduct(input)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})

}

func TestUpdateProduct(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	categoryData := product.Category{
		ID:   1,
		Name: "test",
	}

	var input = product.AddProductInput{
		Name:        "test",
		Description: "test",
		Price:       10000,
		Stock:       10,
		Category:    categoryData,
	}

	mockProduct := product.Product{
		ID:          1,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Category:    input.Category,
	}

	var id = product.GetProductDetailInput{
		ID: 1,
	}
	t.Run("Success update product", func(t *testing.T) {
		repository.On("FindProductByID", id.ID).Return(mockProduct, nil).Once()
		repository.On("Update", mockProduct).Return(mockProduct, nil).Once()
		result, err := usecase.UpdateProduct(id, input)

		assert.NoError(t, err)
		assert.Equal(t, mockProduct, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed update product", func(t *testing.T) {
		repository.On("FindProductByID", id.ID).Return(product.Product{}, errors.New("error")).Once()
		_, err := usecase.UpdateProduct(id, input)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})

}

func TestFindProductByID(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	categoryData := product.Category{
		ID:   1,
		Name: "test",
	}

	var mockProduct = product.Product{
		ID:          1,
		Name:        "test",
		Description: "test",
		Price:       10000,
		Stock:       10,
		Category:    categoryData,
	}

	t.Run("Success find product by id", func(t *testing.T) {
		repository.On("FindProductByID", mockProduct.ID).Return(mockProduct, nil).Once()
		result, err := usecase.FindProductByID(mockProduct.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockProduct, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed find product by id", func(t *testing.T) {
		repository.On("FindProductByID", mockProduct.ID).Return(mockProduct, errors.New("error")).Once()
		_, err := usecase.FindProductByID(mockProduct.ID)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})

}

func TestFindCategoryByID(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var mockCategory = product.Category{
		ID:   1,
		Name: "test",
	}

	t.Run("Success find category by id", func(t *testing.T) {
		repository.On("FindByCategoryByID", mockCategory.ID).Return(mockCategory, nil).Once()
		result, err := usecase.FindCategoryByID(mockCategory.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockCategory, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed find category by id", func(t *testing.T) {
		repository.On("FindByCategoryByID", mockCategory.ID).Return(product.Category{}, errors.New("error")).Once()
		_, err := usecase.FindCategoryByID(mockCategory.ID)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})
}

func TestDeleteProductByID(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var mockProduct = product.Product{
		ID: 1,
	}

	t.Run("Success delete Product by id", func(t *testing.T) {
		repository.On("GetProductByID", mockProduct.ID).Return(mockProduct, nil).Once()
		result, err := usecase.DeleteProductByID(mockProduct.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockProduct, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed delete Product by id", func(t *testing.T) {
		repository.On("GetProductByID", mockProduct.ID).Return(mockProduct, errors.New("error")).Once()
		_, err := usecase.DeleteProductByID(mockProduct.ID)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})
}

func TestDeleteCategoryByID(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var mockCategory = product.Category{
		ID: 1,
	}

	t.Run("Success delete Category by id", func(t *testing.T) {
		repository.On("GetCategoryByID", mockCategory.ID).Return(mockCategory, nil).Once()
		result, err := usecase.DeleteCategoryByID(mockCategory.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockCategory, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed delete Category by id", func(t *testing.T) {
		repository.On("GetCategoryByID", mockCategory.ID).Return(product.Category{}, errors.New("error")).Once()
		_, err := usecase.DeleteCategoryByID(mockCategory.ID)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})
}

func TestCreateCategory(t *testing.T) {

	var repository = mocks.NewRepository(t)
	var usecase = NewUsecase(repository)

	var mockCategory = product.Category{
		Name: "test",
	}

	t.Run("Success create Category", func(t *testing.T) {
		repository.On("SaveCategory", mockCategory).Return(mockCategory, nil).Once()
		result, err := usecase.CreateCategory(mockCategory)
		assert.NoError(t, err)
		assert.Equal(t, mockCategory, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed create mockCategory", func(t *testing.T) {
		repository.On("SaveCategory", mockCategory).Return(product.Category{}, errors.New("error")).Once()
		_, err := usecase.CreateCategory(mockCategory)
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)
	})
}
