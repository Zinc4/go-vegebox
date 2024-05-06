package migration

import (
	"fmt"
	"log"
	"mini-project/user"
	"mini-project/utils/database"
)

func Migration() {
	if err := database.DB.AutoMigrate(&user.User{}, &user.OTP{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")
}
