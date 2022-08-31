package entity

import (
	"log"

	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

type Database struct {
	DatabaseName string
	DatabaseId   string
	InstanceId   string
	LogicalClock int64
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) GetInstanceId() string {
	return d.InstanceId
}

func (d *Database) SetInstanceId(instanceId string) {
	d.InstanceId = instanceId
}

func (d *Database) GetDatabaseName() string {
	return d.DatabaseId
}

func (d *Database) SetDatabaseName(databaseName string) {
	d.DatabaseId = databaseName
}

func (d *Database) GetDatabaseId() string {
	return d.DatabaseId
}

func (d *Database) SetDatabaseId(DatabaseId string) {
	d.DatabaseId = DatabaseId
}

func (d *Database) GetLogicalClock() int64 {
	return d.LogicalClock
}

func (d *Database) SetLogicalClock(logicalClock int64) {
	d.LogicalClock = logicalClock
}

func (d *Database) GetNewDatabaseId(databaseName string) string {
	cid, err := function.GetCidString(databaseName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cid
}
