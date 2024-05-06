package product

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type AddProductInput struct {
	Name        string   `json:"name" form:"name" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	Price       int      `json:"price" form:"price" binding:"required"`
	Stock       int      `json:"stock" form:"stock" binding:"required"`
	Category    Category `json:"category" form:"category" binding:"required"`
}
