package db

//to load test data from csv files to neo4j database

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

//template function to add pincodes,brand and categories
func addMetaData(k, l string) func(m map[string]bool) error {

	f := func(m map[string]bool) error {
		session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer session.Close()

		//cypher query
		q := ""

		for index, _ := range m {
			//to add elements to the cypher query
			q = q + fmt.Sprintf(" (:%s{%s:'%s'}) ,", k, l, index)
		}
		//to remove the last comma
		q = q[0 : len(q)-1]

		//to run the cypher query
		_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				"Create "+q,
				map[string]interface{}{})
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

//to add address pincodes
var AddPincodesTest = addMetaData("Area", "pincode")

//to add product brands
var AddBrandsTest = addMetaData("Brand", "name")

//to add product category
var AddCategoriesTest = addMetaData("Category", "name")

// template function to run cypher queries sent to it
func RunQuery(q string) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			q,
			map[string]interface{}{})
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
