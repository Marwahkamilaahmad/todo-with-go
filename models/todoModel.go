package models

import (
	"github.com/Marwahkamilaahmad/go-fiber-first.git/database"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
)

func GetAllTodos() []entity.Todo {
	var todos []entity.Todo
	result := database.DB.Find(&todos)
	if result.Error != nil {
		panic(result.Error)
	}

	return todos
}

func CreateTodos(entities entity.Todo)bool{

	result := database.DB.Create(&entities)
	if result.Error != nil {
		panic(result.Error)
	}

	return entities.Id > 0
}

func UpdateTodos(entities entity.Todo) bool{

	result := database.DB.Updates(&entities)
	if result.Error != nil {
		panic(result.Error)
	} 

	return true
}


func Detail(id int) entity.Todo {
	var todo entity.Todo
	result := database.DB.First(&todo, id)
	if result.Error != nil {
		panic(result.Error)
	}

	return todo
}

// func Delete(id int) error{

// 	_, err := database.DB.Exec("DELETE FROM todos WHERE id=?", id)
// 	return err
	
// }

func Delete(id int) error {
	var todo entity.Todo
	result := database.DB.Delete(&todo, id)
	if result.Error != nil {
		return result.Error
	}
	
	return nil
}
