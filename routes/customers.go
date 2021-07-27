package routes

import (
	"github.com/gorilla/mux"
	"graphApp/controllers"
)

func CustomerRoutes(c *mux.Router) {
	//route - to add a new customer
	c.HandleFunc("/", controllers.AddCustomer).Methods("POST")
	//route -  to add a relation if customer views a product
	c.HandleFunc("/view", controllers.ViewProduct).Methods("POST")
	//route -  to add a relation if customer wishlists a product
	c.HandleFunc("/wishlist", controllers.WishlistProduct).Methods("POST")
	//route - to add a relation if customer orders a product
	c.HandleFunc("/order", controllers.OrderProduct).Methods("POST")
	//route -  to add a relationship between two customers
	c.HandleFunc("/relation",controllers.AddCustomerRelation).Methods("POST")
}
