package entity

import (
	"log"

	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

// Table  information_schema definition
type Table struct {
	TableId         string
	DatabaseId      string
	TableName       string
	LocalInstanceId string
	LogicalClock    int64
	Version         int64
	Schema          string
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) GetTableId() string {
	return t.TableId
}

func (t *Table) SetTableId(tableId string) {
	t.TableId = tableId
}

func (t *Table) GetDatabaseId() string {
	return t.DatabaseId
}

func (t *Table) SetDatabaseId(databaseId string) {
	t.DatabaseId = databaseId
}

func (t *Table) GetTableName() string {
	return t.TableName
}

func (t *Table) SetTableName(tableName string) {
	t.TableName = tableName
}

func (t *Table) SetLocalInstanceId(localInstanceId string) {
	t.LocalInstanceId = localInstanceId
}

func (t *Table) GetLocalInstanceId() string {
	return t.LocalInstanceId
}

func (t *Table) SetLogicalClock(logicalClock int64) {
	t.LogicalClock = logicalClock
}

func (t *Table) GetLogicalClock() int64 {
	return t.LogicalClock
}

func (t *Table) SetVersion(version int64) {
	t.Version = version
}

func (t *Table) GetVersion() int64 {
	return t.Version
}

func (t *Table) SetSchema(schema string) {
	t.Schema = schema
}

func (t *Table) GetSchema() string {
	return t.Schema
}

func (i *Table) GetNewTableId(tableName string) string {
	cid, err := function.GetCidString(tableName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cid
}
