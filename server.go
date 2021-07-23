package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"graphApp/db"
	"graphApp/global"
	"graphApp/routes"
	"log"
	"net/http"
)

func main() {

	db.Init()
	defer global.Driver.Close()

	r := mux.NewRouter()
	a := r.PathPrefix("/api").Subrouter()
	a.StrictSlash(true)
	a.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		t, err := json.Marshal(map[string]string{"ping": "pong"})

		_, err = w.Write(t)
		if err != nil {
			log.Println(err)
		}
	})

	p := a.PathPrefix("/product").Subrouter()
	c := a.PathPrefix("/customer").Subrouter()
	p.StrictSlash(true)
	c.StrictSlash(true)
	routes.ProductRoutes(p)
	routes.CustomerRoutes(c)

	log.Println("Starting server on port 8080!")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}

//func helloWorld(uri, username, password string) (string, error) {
//	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
//	defer session.Close()
//
//	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
//		result, err := transaction.Run(
//			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
//			map[string]interface{}{"message": "hello, world"})
//		if err != nil {
//			return nil, err
//		}
//
//		if result.Next() {
//			return result.Record().Values[0], nil
//		}
//
//		return nil, result.Err()
//	})
//	if err != nil {
//		return "", err
//	}
//
//	return greeting.(string), nil
//}
