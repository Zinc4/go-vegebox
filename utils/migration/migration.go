package migration

import (
	"fmt"
	"log"
	"mini-project/cart"
	"mini-project/product"
	"mini-project/user"
	"mini-project/utils/database"
)

func Migration() {
	if err := database.DB.AutoMigrate(&user.User{}, &user.OTP{}, &product.Product{}, &product.Category{}, cart.Cart{}, cart.CartItem{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")
}
