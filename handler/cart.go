package handler

import (
	"errors"
	"mini-project/cart"
	"mini-project/helper"
	"mini-project/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type cartHandler struct {
	cartUsecase    cart.Usecase
	productUsecase product.Usecase
}

func NewCartHandler(cartUsecase cart.Usecase, productUsecase product.Usecase) *cartHandler {
	return &cartHandler{cartUsecase, productUsecase}
}

func (h *cartHandler) NewCart(c echo.Context) error {
	newCart, err := h.cartUsecase.CreateCart()
	if err != nil {
		response := helper.ErrorResponse("Error create cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cartt, err := h.cartUsecase.GetCartById(newCart.ID)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	var input cart.RequestBody
	err = c.Bind(&input)
	if err != nil {
		response := helper.ErrorResponse("Error bind input", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	product, err := h.productUsecase.FindProductByID(input.ProductID)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	if product.Stock < input.Quantity {
		response := helper.ErrorResponse("Stock not enough", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cartItem, err := h.cartUsecase.GetCartItemByCartIdAndProductId(cartt.ID, input.ProductID)
	if err != nil {
		response := helper.ErrorResponse("Error get cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	if cartItem.Quantity == 0 {
		cartItem = cart.CartItem{
			CartID:    cartt.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	} else {
		if input.Quantity+cartItem.Quantity > product.Stock {
			reponse := helper.ErrorResponse("Stock not enough", err.Error())
			return c.JSON(http.StatusBadRequest, reponse)
		}
		cartItem.Quantity += input.Quantity
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	}

	if err != nil {
		response := helper.ErrorResponse("Error save cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success add product to cart", h.cartUsecase.GetCartSerializer(cartt))

	return c.JSON(http.StatusOK, response)

}

func (h *cartHandler) AddProductToCart(c echo.Context) error {

	id := c.Param("id")
	cartId, err := strconv.Atoi(id)
	if err != nil {
		response := helper.ErrorResponse("Error convert id", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	cartt, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	var input cart.RequestBody
	err = c.Bind(&input)
	if err != nil {
		response := helper.ErrorResponse("Error bind input", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	product, err := h.productUsecase.FindProductByID(input.ProductID)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	if product.Stock < input.Quantity {
		response := helper.ErrorResponse("Stock not enough", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cartItem, err := h.cartUsecase.GetCartItemByCartIdAndProductId(cartt.ID, input.ProductID)
	if err != nil {
		response := helper.ErrorResponse("Error get cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	if cartItem.Quantity == 0 {
		cartItem = cart.CartItem{
			CartID:    cartt.ID,
			ProductID: input.ProductID,
			Quantity:  input.Quantity,
		}
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	} else {
		if input.Quantity+cartItem.Quantity > product.Stock {
			response := helper.ErrorResponse("Stock not enough", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		cartItem.Quantity += input.Quantity
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	}

	if err != nil {
		response := helper.ErrorResponse("Error save cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success add product to cart", h.cartUsecase.GetCartSerializer(cartt))

	return c.JSON(http.StatusOK, response)

}

func (h *cartHandler) UpdateCartItem(c echo.Context) error {
	cartId, err := strconv.Atoi(c.Param("cart_id"))
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusNotFound, response)
	}
	cartt, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	var input cart.UpdateCartQuantity
	err = c.Bind(&input)
	if err != nil {
		response := helper.ErrorResponse("Error bind input", err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	product, err := h.productUsecase.FindProductByID(productID)
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	cartItem, err := h.cartUsecase.GetCartItemByCartIdAndProductId(cartt.ID, productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		} else {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	if cartItem.Quantity == 0 {
		cartItem = cart.CartItem{
			CartID:    cartt.ID,
			ProductID: productID,
			Quantity:  input.Quantity,
		}
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	} else {
		if input.Quantity+cartItem.Quantity > product.Stock {
			response := helper.ErrorResponse("Stock not enough", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
		cartItem.Quantity = input.Quantity
		cartItem, err = h.cartUsecase.SaveCartItem(cartItem)
	}

	if err != nil {
		response := helper.ErrorResponse("Error save cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success update cart", h.cartUsecase.GetCartSerializer(cartt))

	return c.JSON(http.StatusOK, response)

}

func (h *cartHandler) GetCart(c echo.Context) error {

	id := c.Param("id")
	cartId, err := strconv.Atoi(id)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	cart, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success get cart", h.cartUsecase.GetCartSerializer(cart))

	return c.JSON(http.StatusOK, response)
}

func (h *cartHandler) GetCarts(c echo.Context) error {

	carts, err := h.cartUsecase.GetCarts()
	if err != nil {
		response := helper.ErrorResponse("Error get carts", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseWithData("Success get carts", carts)

	return c.JSON(http.StatusOK, response)
}

func (h *cartHandler) DeleteCart(c echo.Context) error {
	id := c.Param("id")
	cartId, err := strconv.Atoi(id)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	cart, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cartItems, err := h.cartUsecase.GetCartItemByCartId(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart items", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	for _, item := range cartItems {
		err := h.cartUsecase.DeleteCartItemByCartIdAndProductId(cartId, item.ProductID)
		if err != nil {
			response := helper.ErrorResponse("Error delete cart item", err.Error())
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	err = h.cartUsecase.DeleteCartById(cart.ID)
	if err != nil {
		response := helper.ErrorResponse("Error delete cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("success delete cart"))
}

func (h *cartHandler) DeleteProductFromCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("cart_id"))
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		response := helper.ErrorResponse("Error get product", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.cartUsecase.DeleteCartItemByCartIdAndProductId(cartID, productID)
	if err != nil {
		response := helper.ErrorResponse("Error delete cart item", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("successfully deleted product from cart"))
}

func (h *cartHandler) CheckOut(c echo.Context) error {

	id := c.Param("id")
	cartId, err := strconv.Atoi(id)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cart, err := h.cartUsecase.GetCartById(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	cartItems, err := h.cartUsecase.GetCartItemByCartId(cartId)
	if err != nil {
		response := helper.ErrorResponse("Error get cart items", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	totalPrice := int64(0)
	for _, item := range cartItems {
		totalPrice += int64(item.Product.Price) * int64(item.Quantity)
	}

	return c.JSON(http.StatusOK, helper.ResponseWithData("successfully checked out", map[string]interface{}{"cart": cart, "total_price": totalPrice}))

}
