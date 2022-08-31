package repository

import (
	PO "github.com/Rock-liyi/p2pdb/domain/log/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

var DB *orm.CreateDBFactory = new(orm.CreateDBFactory)

func init() {
	//settings  use  the default  information database
	DB.SetIsInternalStore(true)
}

func RemoveDatabase(databasePath string) {
	log.Info("databasePath is" + databasePath)
	err := function.RemoveFile(databasePath)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetDatabasePath(databaseName string) string {
	DB.InitDB()
	//DB.SetDatabaseName(databaseName)
	return DB.GetDatabasePath(databaseName)
}

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

func InsertNodeLog(Po *PO.Node, tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewNode()
	newTable.SetTabName(tableName)
	db.InsertIgnore(Po)
	return true
}

func InsertDataLog(Po *PO.Data, tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewData()
	newTable.SetTabName(tableName)
	db.InsertIgnore(Po)
	return true
}

func InsertLinkLog(Po *PO.Link, tableName string, databaseName string) bool {
	DB.SetDatabaseName(databaseName)
	db := DB.InitDB()
	var newTable = PO.NewLink()
	newTable.SetTabName(tableName)
	db.InsertIgnore(Po)
	return true
}
