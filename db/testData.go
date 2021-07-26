package db

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

func AddPincodesTest(p map[string]bool) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer session.Close()

	q := ""

	for index, _ := range p {
		q = q + fmt.Sprintf(" (:Area{pincode:%s}) ,", index)
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



func AddBrandsTest(b map[string]bool) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	q := ""

	for index, _ := range b {
		q = q + fmt.Sprintf("(:Brand{name:'%s'}) ,", index)
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

func AddCategoriesTest(c map[string]bool) error {
	session := global.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	q := ""

	for index, _ := range c {
		q = q + fmt.Sprintf("(:Category{name:'%s'}) ,", index)
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
