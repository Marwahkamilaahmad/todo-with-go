package models

import (

	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
)

func GetAllTodos() (entity.Todo) {
	var todos entity.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		panic(result.Error)
	}

	return todos
}