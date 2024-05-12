package cart

type RequestBody struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UpdateCartQuantity struct {
	Quantity int `json:"quantity"`
}
