package migration

import (
	"fmt"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("success to migrate")
}