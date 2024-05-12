package product

type Usecase interface {
	FindProducts() ([]Product, error)
	FindProductByID(ID int) (Product, error)
	FindAllCategory() ([]Category, error)
	FindProductByCategory(categoryID int) ([]Product, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) FindProducts() ([]Product, error) {
	products, err := u.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (u *usecase) FindProductByID(ID int) (Product, error) {
	product, err := u.repository.FindByID(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (u *usecase) FindAllCategory() ([]Category, error) {
	categories, err := u.repository.FindAllCategory()
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (u *usecase) FindProductByCategory(categoryID int) ([]Product, error) {
	products, err := u.repository.FindByCategory(categoryID)
	if err != nil {
		return products, err
	}
	return products, nil
}
