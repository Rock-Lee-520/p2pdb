package repository

import (
	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
)

var DB *orm.CreateDBFactory = new(orm.CreateDBFactory)

func init() {
	//settings  use  the default  information database
	DB.SetIsDBInformation(true)
}

func CreateDatabaseInformationTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.DatabaseInfomation{})
	return true
}

func CreateTableInformationTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.TableInfomation{})
	return true
}
