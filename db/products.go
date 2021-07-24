package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

func AddP(name string, sku string, id string, price float32, description string, brand string, category string) (string, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (b:Brand{name:$brand}) "+
				"MERGE (c:category{name:$category}) "+
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

func ViewP(cId, pId, date string) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (c:Customer{internal_id:$cId}),(p:Product{internal_id:$pId}) "+
				"MERGE (c)-[:View{date:$date}]->(p)",
			map[string]interface{}{"cId": cId, "pId": pId, "date": date})

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

func WishlistP(cId, pId, date string) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (c:Customer{internal_id:$cId}),(p:Product{internal_id:$pId}) "+
				"MERGE (c)-[:Wishlist{date:$date}]->(p) "+
				"MATCH (c)-[r:View]->(p) "+
				"DELETE r",

			map[string]interface{}{"cId": cId, "pId": pId, "date": date})

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

func OrderP(cId, pId, date string) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (c:Customer{internal_id:$cId}),(p:Product{internal_id:$pId}) "+
				"MERGE (c)-[:Order{date:$date}]->(p) "+
				"PARTIAL MATCH (c)-[r:View]->(p),(c)-[o:Wishlist]->(p) "+
				"DELETE r , p",

			map[string]interface{}{"cId": cId, "pId": pId, "date": date})

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
