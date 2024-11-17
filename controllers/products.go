package controllers

import (
	"alura-web-1/db"
	"alura-web-1/models"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := db.FindAll()
	temp.ExecuteTemplate(w, "index", products)
}

func Novo(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "novo", nil)
}

func Salvar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")

		priceConverted, err := strconv.ParseFloat(price, 32)

		if err != nil {
			panic(err)
		}

		productToSave := models.Product{
			Name:        name,
			Description: description,
			Price:       float32(priceConverted),
		}

		if id == "" {
			db.Save(productToSave)
		} else {
			idConverted, err := strconv.Atoi(id)

			if err != nil {
				panic(err)
			}
			productToSave.Id = idConverted
			db.Update(productToSave)
		}
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.URL.Query().Get("productId"))

	if err != nil {
		panic(err)
	}

	productFound := db.FindById(productId)

	temp.ExecuteTemplate(w, "editar", productFound)
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.URL.Query().Get("productId"))

	msg := "PRODUCT ID DELETE " + strconv.Itoa(productId)
	slog.Info(msg)

	if err != nil {
		panic(err)
	}

	db.DeleteById(productId)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
