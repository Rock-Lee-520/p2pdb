package entity

import (
	"log"

	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

type DatabaseEntity struct {
	DatabaseName string
	DatabaseId   string
	InstanceId   string
}

func NewDatabaseEntity() *DatabaseEntity {
	return &DatabaseEntity{}
}

func (i *DatabaseEntity) GetInstanceId() string {
	return i.InstanceId
}

func (i *DatabaseEntity) SetInstanceId(instanceId string) {
	i.InstanceId = instanceId
}

func (i *DatabaseEntity) GetDatabaseName() string {
	return i.DatabaseId
}

func (i *DatabaseEntity) SetDatabaseName(databaseName string) {
	i.DatabaseId = databaseName
}

func (i *DatabaseEntity) GetDatabaseId() string {
	return i.DatabaseId
}

func (i *DatabaseEntity) SetDatabaseId(DatabaseId string) {
	i.DatabaseId = DatabaseId
}

func GetNewDatabaseId(databaseName string) string {
	cid, err := function.GetCidString(databaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cid
}
