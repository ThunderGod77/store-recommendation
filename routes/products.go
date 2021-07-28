package routes

import (
	"github.com/gorilla/mux"
	"graphApp/controllers"
)

func ProductRoutes(p *mux.Router) {
	// route - to add a product
	p.HandleFunc("/", controllers.AddProduct).Methods("POST")
	// route - to delete a product
	p.HandleFunc("/{id}", controllers.DeleteProduct).Methods("DELETE")
	//route - to add a brand
	p.HandleFunc("/brand", controllers.AddBrand).Methods("POST")
	//route - to add product category
	p.HandleFunc("/category", controllers.AddCategory).Methods("POST")
	//route -  to get recommendations
	p.HandleFunc("/recommend/{pId}", controllers.Recommend).Methods("GET")
}
