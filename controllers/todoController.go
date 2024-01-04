package controllers

import (
	"net/http"
	"text/template"
	// "time"

	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
	// "github.com/Marwahkamilaahmad/go-fiber-first.git/entity/migration"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/models"
	// "github.com/gofiber/fiber/v2"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()
	data := map[string]interface{}{
		"todos" : todos,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// temp.Execute(w, data)
	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/create.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var todos entity.Todo

		todos.Judul = r.FormValue("judul")
		todos.Hari = r.FormValue("hari")
		todos.Waktu = r.FormValue("waktu")
		todos.Keterangan = r.FormValue("keterangan")
		// todos.CreatedAt = time.Now()
		// todos.UpdatedAt = time.Now()

		if ok := models.CreateTodos(todos); !ok{
			temp, _ := template.ParseFiles("views/create.html")
			temp.Execute(w, nil)

		}
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {

	return 
}

func Delete(w http.ResponseWriter, r *http.Request) {
	return 

}