package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

func AddC(name string, internalId string, pincode int, email string) (string, error) {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MERGE (a:Area{pincode:$pincode}) "+
				"CREATE (c:Customer)-[:In]->(a) "+
				"SET "+
				"c.name = $name,"+
				"c.internal_id = $id,"+
				"c.email = $email "+
				"RETURN c.id",
			map[string]interface{}{"name": name, "email": email, "id": internalId, "pincode": pincode})

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
