package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

type pName struct {
	Names []string `json:"names"`
}

//other things bought by people who bought this product
func AlsoBought(productId, customerId string) (interface{}, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (p:Product{internal_id:$pId})<-[:Order]-(c:Customer)-[:Order]->(pr:Product) where not (c.internal_id=$cId or pr.internal_id=$pId) return pr.name AS name",
			map[string]interface{}{"cId": customerId, "pId": productId})

		if err != nil {
			return nil, err
		}
		var pr []string
		for result.Next() {
			record := result.Record()
			name, ok := record.Get("name")
			if !ok {
				name = ""
			}
			pr = append(pr, fmt.Sprintf(name.(string)))
		}

		return pName{Names: pr}, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return result, nil

}

//products of same category
func SameCategory(pId, cId string) (interface{}, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (cr:Customer{internal_id:$cId})-[:Order]->(p:Product{internal_id:$pId})-[:Belongs]->(c:Category)<-[:Belongs]->(pr:Product) where not (cr)-[:Order]->(pr) return pr.name as name",
			map[string]interface{}{"cId": cId, "pId": pId})

		if err != nil {
			return nil, err
		}
		var pr []string
		for result.Next() {
			record := result.Record()
			name, ok := record.Get("name")
			if !ok {
				name = ""
			}
			pr = append(pr, fmt.Sprintf(name.(string)))
		}

		return pName{Names: pr}, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

//products of same brand
func SameBrand(pId, cId string) (interface{}, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (cr:Customer{internal_id:$cId})-[:Order]->(p:Product{internal_id:$pId})-[:By]->(c:Brand)<-[:By]->(pr:Product) where not (cr)-[:Order]->(pr) return pr.name as name",
			map[string]interface{}{"cId": cId, "pId": pId})

		if err != nil {
			return nil, err
		}
		var pr []string
		for result.Next() {
			record := result.Record()
			name, ok := record.Get("name")
			if !ok {
				name = ""
			}
			pr = append(pr, fmt.Sprintf(name.(string)))
		}


		return pName{Names: pr}, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}
