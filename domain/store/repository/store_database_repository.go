package repository

import (
	"time"

	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func FindDatabaseInformation(DatabaseId string) *PO.DatabaseInfomation {
	db = InitDB()
	var database = &PO.DatabaseInfomation{}
	db.Where(database.GetDatabasePrimary()+" = ?", DatabaseId).First(database)

	return database
}

func UpdateDatabaseInformation(databasePO *PO.DatabaseInfomation) bool {
	db = InitDB()
	databasePO.UpdatedAt = time.Now()
	//err := db.Where(databasePO.GetDatabasePrimary()+" = ?", databasePO.DatabaseId).Updates(databasePO)
	//if err != nil {
	//	log.Error(err)
	//	return false
	//}
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
