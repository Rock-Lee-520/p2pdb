package repository

import (
	"time"

	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func FindTableInformation(DatabaseId string) *PO.TableInfomation {
	db = InitDB()
	var table = &PO.TableInfomation{}
	db.Where(table.GetTablePrimary()+" = ?", DatabaseId).First(table)

	return table
}

func UpdateTableInformation(tablePO *PO.TableInfomation) bool {
	db = InitDB()
	tablePO.UpdatedAt = time.Now()
	err := db.Where(tablePO.GetTablePrimary()+" = ?", tablePO.TableId).Updates(tablePO)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func InsertTableInformation(tablePO *PO.TableInfomation) bool {
	db = InitDB()
	err := db.InsertIgnore(tablePO)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func CreateTableInformationTable() bool {
	db = InitDB()
	db.Migrator().CreateTable(&PO.TableInfomation{})
	return true
}
