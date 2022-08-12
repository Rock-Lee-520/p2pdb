package repository

import (
	"time"

	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

var ORM *orm.CreateDBFactory = new(orm.CreateDBFactory)
var db orm.DBconnect

func InitDB() orm.DBconnect {
	//settings  use  the default  information database
	ORM.SetIsInternalStore(true)
	db = ORM.InitDB()
	return db
}

func FindDatabaseInformationTable(DatabaseId string) *PO.DatabaseInfomation {
	db = InitDB()
	var database = &PO.DatabaseInfomation{}
	db.Where(database.GetDatabasePrimaryId()+" = ?", DatabaseId).First(database)

	return database
}

func UpdateDatabaseInformation(databasePO *PO.DatabaseInfomation) bool {
	db = InitDB()
	databasePO.UpdatedAt = time.Now()
	err := db.Where(databasePO.GetDatabasePrimaryId()+" = ?", databasePO.DatabaseId).Updates(databasePO)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func InsertDatabaseInformation(databasePO *PO.DatabaseInfomation) bool {
	db = InitDB()
	err := db.InsertIgnore(databasePO)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func CreateDatabaseInformationTable() bool {
	db = InitDB()
	db.Migrator().CreateTable(&PO.DatabaseInfomation{})
	return true
}

func CreateTableInformationTable() bool {
	db = InitDB()
	db.Migrator().CreateTable(&PO.TableInfomation{})
	return true
}
