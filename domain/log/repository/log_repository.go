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

// func CreateNodeTable() bool {
// 	db := DB.InitDB()
// 	db.Migrator().CreateTable(&PO.Node{})
// 	return true
// }

// func CreateDataTable() bool {
// 	db := DB.InitDB()
// 	db.Migrator().CreateTable(&PO.Data{})
// 	return true
// }

// func CreateLinkTable() bool {
// 	db := DB.InitDB()
// 	db.Migrator().CreateTable(&PO.Link{})
// 	return true
// }

func CreateNodeTable(tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewNode()
	newTable.SetTabName(tableName)
	db.Migrator().AutoMigrate(newTable)
	return true
}

func CreateDataTable(tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewData()
	newTable.SetTabName(tableName)
	db.Migrator().AutoMigrate(newTable)
	return true
}

func CreateLinkTable(tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewLink()
	newTable.SetTabName(tableName)
	db.Migrator().AutoMigrate(newTable)
	return true
}
