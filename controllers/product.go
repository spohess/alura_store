package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/spohess/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic("ID inválido")
	}
	product := models.GetProductById(id)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		panic("Valor inválido")
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		panic("Quantidade inválido")
	}

	models.CreateProduct(name, description, price, quantity)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic("ID inválido")
	}

	models.DeleteProduct(id)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
