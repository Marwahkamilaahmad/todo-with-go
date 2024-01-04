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

func CreateTodos(entities entity.Todo)bool{
	// result, err := database.DB.Exec(
	// 	`INSERT INTO todos(judul, hari, waktu, keterangan, created_at, updated_at)
	// 	VALUE(?,?,?,?,?,?)`,
	// 	entities.Judul, entities.Hari, entities.Waktu, entities.CreatedAt, entities.UpdatedAt,
	// )

	// if err != nil {
	// 	panic(err)
	// }

	result := database.DB.Create(&entities)
	if result.Error != nil {
		panic(result.Error)
	}

	return entities.Id > 0
}

func Edit(){
	
}