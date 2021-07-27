package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"graphApp/controllers/data"
	"graphApp/db"
	"graphApp/global"
	"graphApp/routes"
	"log"
	"net/http"
)

func main() {

	//initializing the neo4j driver
	db.Init()
	defer global.Driver.Close()



	r := mux.NewRouter()
	a := r.PathPrefix("/api").Subrouter()
	a.StrictSlash(true)

	// setting up test route
	a.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		t, err := json.Marshal(map[string]string{"ping": "pong"})

		_, err = w.Write(t)
		if err != nil {
			log.Println(err)
		}
	})

	//route to reset database and load test data
	a.HandleFunc("/load-data", data.LoadTestData).Methods("GET")

	//subrouter related to product
	p := a.PathPrefix("/product").Subrouter()
	////subrouter related to customer
	c := a.PathPrefix("/customer").Subrouter()
	p.StrictSlash(true)
	c.StrictSlash(true)

	//registers all the routes
	routes.ProductRoutes(p)
	routes.CustomerRoutes(c)

	log.Println("Starting server on port 8080!")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}

