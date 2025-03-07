package controllers

import (
	"html/template"
	"net/http"

	"github.com/spohess/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
