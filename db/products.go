package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)
//to add product
func AddP(name string, sku string, id string, price float32, description string, brand string, category string) (string, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	//actual cypher query
	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (b:Brand{name:$brand}) "+
				"MERGE (c:Category{name:$category}) "+
				"CREATE (a:Product)-[:By]->(b),(a)-[:Belongs]->(c) "+
				"SET "+
				"a.name = $name,"+
				"a.sku = $sku,"+
				"a.internal_id = $id,"+
				"a.price = $price,"+
				"a.description = $description "+
				"RETURN a.id",
			map[string]interface{}{"name": name, "sku": sku, "id": id, "price": price, "description": description, "brand": brand, "category": category})

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

//to delete products
func DeleteP(id string) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH(n:Product{internal_id:$id}) DETACH DELETE n",
			map[string]interface{}{"id": id})

		if err != nil {
			return nil, err
		}
		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return nil

}

//template function to add metadata related to products
func addMData(key string) func(string) error {
	f := func(name string) error {
		session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer session.Close()

		_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				"MERGE (:$key{name:$name})",
				map[string]interface{}{"key": key, "name": name})

			if err != nil {
				return nil, err
			}
			return nil, result.Err()
		})

		if err != nil {
			return err
		}

		return nil
	}
	return f
}
// to add brand nodes
var AddBrand = addMData("Brand")
//to add category nodes
var AddCategory = addMData("Category")

//template function to add customer product relations
func cpAction(key string) func(cId, pId, date string) error {
	f := func(cId, pId, date string) error {
		session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer session.Close()

		_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				"MATCH (c:Customer{internal_id:$cId}),(p:Product{internal_id:$pId}) "+
					"MERGE (c)-[:$action{date:$date}]->(p)",
				map[string]interface{}{"cId": cId, "pId": pId, "date": date, "action": key})

			if err != nil {
				return nil, err
			}
			return nil, result.Err()
		})

		if err != nil {
			return err
		}

		return nil
	}
	return f
}
//to deal if customer views the product
var ViewP = cpAction("View")
//to deal with if customer wishlists the product
var WishlistP = cpAction("Wishlist")
//to deal with if customer orders a product
var OrderP = cpAction("Order")
