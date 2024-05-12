package mocks

import (
	"errors"
	"mini-project/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindProducts(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = product.NewUsecase(repository)

	var products = []product.Product{
		{
			ID:          1,
			Name:        "test",
			Description: "test",
			Price:       10000,
			Stock:       10,
		},
		{
			ID:          2,
			Name:        "test",
			Description: "test",
			Price:       10000,
			Stock:       10,
		},
	}

	t.Run("Success find products", func(t *testing.T) {

		repository.On("FindAll").Return(products, nil).Once()

		result, err := usecase.FindProducts()
		assert.NoError(t, err)
		assert.Equal(t, products, result)
		repository.AssertExpectations(t)

	})

	t.Run("Failed find products", func(t *testing.T) {

		repository.On("FindAll").Return(products, errors.New("error")).Once()

		result, err := usecase.FindProducts()
		assert.Error(t, err)
		assert.Equal(t, "error", err.Error())
		repository.AssertExpectations(t)

		assert.Equal(t, products, result)
	})

}

func TestFindProductByID(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = product.NewUsecase(repository)

	var prodct = product.Product{
		ID:          1,
		Name:        "test",
		Description: "test",
		Price:       10000,
		Stock:       10,
	}

	t.Run("Success find product by id", func(t *testing.T) {

		repository.On("FindByID", prodct.ID).Return(prodct, nil).Once()

		result, err := usecase.FindProductByID(prodct.ID)
		assert.NoError(t, err)
		assert.Equal(t, prodct, result)

		repository.AssertExpectations(t)

	})

	t.Run("Failed find product by id", func(t *testing.T) {

		repository.On("FindByID", prodct.ID).Return(prodct, errors.New("error")).Once()

		_, err := usecase.FindProductByID(prodct.ID)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		repository.AssertExpectations(t)
	})

}

func TestFindAllCategory(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = product.NewUsecase(repository)

	var categories = []product.Category{
		{
			ID:   1,
			Name: "Sayuran",
		},
		{
			ID:   2,
			Name: "Buah",
		},
	}

	t.Run("Success find all category", func(t *testing.T) {

		repository.On("FindAllCategory").Return(categories, nil).Once()

		result, err := usecase.FindAllCategory()

		assert.NoError(t, err)

		assert.Equal(t, categories, result)

		repository.AssertExpectations(t)

	})

	t.Run("Failed find all category", func(t *testing.T) {

		repository.On("FindAllCategory").Return(categories, errors.New("error")).Once()

		_, err := usecase.FindAllCategory()

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		repository.AssertExpectations(t)
	})
}

func TestFindProductByCategory(t *testing.T) {

	var repository = NewRepository(t)
	var usecase = product.NewUsecase(repository)

	var products = []product.Product{
		{
			ID:          1,
			Name:        "test",
			Description: "test",
			Price:       10000,
			Stock:       10,
		},
		{
			ID:          2,
			Name:        "test",
			Description: "test",
			Price:       10000,
			Stock:       10,
		},
	}

	t.Run("Success find product by category", func(t *testing.T) {

		repository.On("FindByCategory", 1).Return(products, nil).Once()

		result, err := usecase.FindProductByCategory(1)

		assert.NoError(t, err)

		assert.Equal(t, products, result)

		repository.AssertExpectations(t)
	})

	t.Run("Failed find product by category", func(t *testing.T) {

		repository.On("FindByCategory", 1).Return(products, errors.New("error")).Once()

		_, err := usecase.FindProductByCategory(1)

		assert.Error(t, err)

		assert.Equal(t, "error", err.Error())

		repository.AssertExpectations(t)

	})
}
