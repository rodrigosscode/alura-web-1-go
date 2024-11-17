package main

import (
	"alura-web-1/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":2323", nil)
}
