package repository

import (
	PO "github.com/Rock-liyi/p2pdb/domain/log/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
)

var DB *orm.CreateDBFactory = new(orm.CreateDBFactory)

func init() {
	//settings  use  the default  information database
	DB.SetIsInternalStore(true)
}

func CreateNodeTable(tableName string) bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Node{})
	return true
}

func CreateDataTable(tableName string) bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Data{})
	return true
}

func CreateLinkTable(tableName string) bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.Data{})
	return true
}
