package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	// "time"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/entity"
	// "github.com/Marwahkamilaahmad/go-fiber-first.git/entity/migration"
	"github.com/Marwahkamilaahmad/go-fiber-first.git/models"
	// "github.com/gofiber/fiber/v2"
)

var BASE_URL = "10.33.35.56/todos"

func IndexConsume(w http.ResponseWriter, r *http.Request) {
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

func AddConsume(w http.ResponseWriter, r *http.Request) {

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

func EditConsume(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		todo:= models.Detail(id)
		data := map[string]interface{}{
			"todo": todo,
		}

		err = temp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} 

	if r.Method == "POST" {
		var todos entity.Todo

		id_form := r.FormValue("id")
		idString, err := strconv.Atoi(id_form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
		todos.Id = uint(idString)

		todos.Hari = r.FormValue("hari")
		todos.Judul = r.FormValue("judul")
		todos.Keterangan = r.FormValue("keterangan")
		todos.Waktu = r.FormValue("waktu")


		ok := models.UpdateTodos(todos)
		if !ok {
			http.Error(w, "Failed to update", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}

}


func DeleteConsume(w http.ResponseWriter, r *http.Request) {

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err:= models.Delete(id); err != nil {
			panic(err)
		}
		http.Redirect(w,r, "/todos", http.StatusSeeOther)

}

