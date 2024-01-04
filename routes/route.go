package routes

import (

	"net/http"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/controllers"
)

func RouterApp(){
	http.HandleFunc("/todos", controllers.Index)
	http.HandleFunc("/todos/add", controllers.Add)
	http.HandleFunc("/todos/edit", controllers.Edit)
	http.HandleFunc("/todos/delete", controllers.Delete)
}