package routes

import (
	"github.com/gorilla/mux"
	"graphApp/controllers"
)

func ProductRoutes(p *mux.Router) {
	p.HandleFunc("/", controllers.AddProduct).Methods("POST")
	p.HandleFunc("/{id}", controllers.DeleteProduct).Methods("DELETE")
}
