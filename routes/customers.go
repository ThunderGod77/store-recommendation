package routes

import (
	"github.com/gorilla/mux"
	"graphApp/controllers"
)

func CustomerRoutes(c *mux.Router) {
	c.HandleFunc("/", controllers.AddCustomer).Methods("POST")
	c.HandleFunc("/view", controllers.ViewProduct).Methods("POST")
	c.HandleFunc("/wishlist", controllers.WishlistProduct).Methods("POST")
	c.HandleFunc("/order", controllers.OrderProduct).Methods("POST")
	c.HandleFunc("/relation",controllers.AddCustomerRelation).Methods("POST")
}
