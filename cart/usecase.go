package cart

type Usecase interface {
	CreateCart() (Cart, error)
	GetCartItemByCartId(cartID int) ([]CartItem, error)
	SaveCartItem(cartItem CartItem) (CartItem, error)
	DeleteCartById(cartID int) error
	DeleteCartItemByCartIdAndProductId(cartID int, productID int) error
	GetCartById(id int) (Cart, error)
	GetCartItemByCartIdAndProductId(cartID int, productID int) (CartItem, error)
	GetCartSerializer(cart Cart) interface{}
	GetCarts() ([]Cart, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{repository}
}

func (u *usecase) CreateCart() (Cart, error) {
	cart := Cart{}
	result, err := u.repository.Create(cart)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *usecase) GetCartItemByCartId(cartID int) ([]CartItem, error) {
	cartItems, err := u.repository.GetCartItemByCartId(cartID)
	if err != nil {
		return cartItems, err
	}
	return cartItems, nil
}

func (u *usecase) SaveCartItem(cartItem CartItem) (CartItem, error) {
	result, err := u.repository.SaveCartItem(cartItem)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *usecase) DeleteCartById(cartID int) error {
	err := u.repository.DeleteById(cartID)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) GetCartById(id int) (Cart, error) {
	cart, err := u.repository.GetById(id)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (u *usecase) DeleteCartItemByCartIdAndProductId(cartID int, productID int) error {
	err := u.repository.DeleteCartItemByCartIdAndProductId(cartID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) DeleteProductFromCart(cartID int, productID int) error {
	err := u.repository.DeleteProductByCartIdAndProductId(cartID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) GetCartItemByCartIdAndProductId(cartID int, productID int) (CartItem, error) {
	cartItem, err := u.repository.GetCartItemByCartIdAndProductId(cartID, productID)
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

func (u *usecase) GetCartSerializer(cart Cart) interface{} {
	items, _ := u.GetCartItemByCartId(cart.ID)
	res_items := make([]CartItemFormatter, len(items))
	for i, item := range items {
		res_items[i].CreatedAt = item.CreatedAt
		res_items[i].Product = item.Product
		res_items[i].Quantity = item.Quantity
	}

	cart_ser := CartsFormatter{
		ID:        cart.ID,
		Items:     res_items,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
		DeletedAt: cart.DeletedAt,
	}

	return cart_ser

}

func (u *usecase) GetCarts() ([]Cart, error) {
	carts, err := u.repository.GetCarts()
	if err != nil {
		return carts, err
	}
	return carts, nil
}
