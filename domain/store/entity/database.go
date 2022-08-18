package entity

import (
	"log"

	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

type Database struct {
	DatabaseName string
	DatabaseId   string
	InstanceId   string
}

func NewDatabase() *Database {
	return &Database{}
}

func (i *Database) GetInstanceId() string {
	return i.InstanceId
}

func (i *Database) SetInstanceId(instanceId string) {
	i.InstanceId = instanceId
}

func (i *Database) GetDatabaseName() string {
	return i.DatabaseId
}

func (i *Database) SetDatabaseName(databaseName string) {
	i.DatabaseId = databaseName
}

func (i *Database) GetDatabaseId() string {
	return i.DatabaseId
}

func (i *Database) SetDatabaseId(DatabaseId string) {
	i.DatabaseId = DatabaseId
}

func GetNewDatabaseId(databaseName string) string {
	cid, err := function.GetCidString(databaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cid
}
