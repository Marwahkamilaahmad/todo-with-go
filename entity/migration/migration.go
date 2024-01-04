package migration

import (
	"fmt"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
)

func RunMigrate() {
	errUser := database.DB.AutoMigrate(&entity.User{})
	if errUser != nil {
		panic(errUser)
	}

	errTodo := database.DB.AutoMigrate(&entity.Todo{})
	if errTodo != nil {
		panic(errTodo)
	}

	fmt.Println("Success migrating User and Todo tables")
}