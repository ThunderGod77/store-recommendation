package routes

import (
	"github.com/gorilla/mux"
	"graphApp/controllers"
)

func CustomerRoutes(c *mux.Router) {
	c.HandleFunc("/", controllers.AddCustomer).Methods("POST")
	c.HandleFunc("/view", controllers.ViewProduct).Methods("POST")
	c.HandleFunc("/wishlist", controllers.WishlistProduct).Methods("POST")
}
