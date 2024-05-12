package cart

import "gorm.io/gorm"

type Repository interface {
	Create(cart Cart) (Cart, error)
	GetCartItemByCartId(cartID int) ([]CartItem, error)
	GetById(id int) (Cart, error)
	DeleteById(id int) error
	DeleteProductByCartIdAndProductId(cartID int, productID int) error
	DeleteCartItemByCartIdAndProductId(cartID int, productID int) error
	SaveCartItem(cartItem CartItem) (CartItem, error)
	GetCartItemByCartIdAndProductId(cartID int, productID int) (CartItem, error)
	GetCarts() ([]Cart, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(cart Cart) (Cart, error) {
	err := r.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) GetCartItemByCartId(cartID int) ([]CartItem, error) {
	var cartItems []CartItem
	err := r.db.Preload("Product").Preload("Product.Category").Preload("Cart").Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		return cartItems, err
	}
	return cartItems, nil
}

func (r *repository) GetById(id int) (Cart, error) {
	var cart Cart
	err := r.db.Where("id = ?", id).First(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) DeleteById(id int) error {
	err := r.db.Where("id = ?", id).Delete(&Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteCartItemByCartIdAndProductId(cartID int, productID int) error {
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&CartItem{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteProductByCartIdAndProductId(cartID int, productID int) error {
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&CartItem{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) SaveCartItem(cartItem CartItem) (CartItem, error) {
	err := r.db.Save(&cartItem).Error
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

func (r *repository) GetCartItemByCartIdAndProductId(cartID int, productID int) (CartItem, error) {
	var cartItem CartItem
	err := r.db.Preload("Product").Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartItem).Error
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

func (r *repository) GetCarts() ([]Cart, error) {
	var carts []Cart
	err := r.db.Find(&carts).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}
