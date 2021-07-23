package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"graphApp/global"
)

var err error

func Init() {
	global.Driver, err = neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "recommend", ""))
	if err != nil {
		panic(err)
	}

}
