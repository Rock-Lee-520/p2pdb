package service

import (
	"testing"

	commonEntity "github.com/Rock-liyi/p2pdb/domain/common/entity"
)

func TestAddTableRecord(t *testing.T) {
	var data = commonEntity.NewData()
	data.SetSQLStatement(" INSERT INTO p2pdb.test6 (name2,email2,id) VALUES('123','123',123)")
	data.SetDatabaseName("p2pdb")
	data.SetDMLType("insert")
	data.SetTableName("test3")

	AddTableRecord(*data)
}
