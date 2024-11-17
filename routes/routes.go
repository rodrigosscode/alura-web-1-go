package routes

import (
	"alura-web-1/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novo", controllers.Novo)
	http.HandleFunc("/salvar", controllers.Salvar)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/delete", controllers.Deletar)
}
