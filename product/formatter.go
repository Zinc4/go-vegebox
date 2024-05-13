package product

type productFormatter struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
}

func FormatProduct(product Product) productFormatter {

	return productFormatter{

		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func FormatProducts(products []Product) []productFormatter {

	var productFormatters []productFormatter

	for _, product := range products {

		productFormatters = append(productFormatters, FormatProduct(product))
	}

	return productFormatters

}
