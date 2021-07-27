package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

func AddMetaData(k, l string) func(m map[string]bool) error {

	f := func(m map[string]bool) error {
		session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		defer session.Close()

		q := ""

		for index, _ := range m {
			q = q + fmt.Sprintf(" (:%s{%s:'%s'}) ,", k, l, index)
		}
		q = q[0 : len(q)-1]

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

var AddPincodesTest = AddMetaData("Area", "pincode")
var AddBrandsTest = AddMetaData("Brand", "name")
var AddCategoriesTest = AddMetaData("Category", "name")

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
