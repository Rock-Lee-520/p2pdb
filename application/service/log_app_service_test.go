package service

import (
	"testing"

	commonEntity "github.com/Rock-liyi/p2pdb/domain/common/entity"
)

func TestCreateStoreLog(t *testing.T) {
	//CreateStoreLog()
}

func TestInsertStoreLog(t *testing.T) {
	var data = commonEntity.NewData()
	data.SQLStatement = " INSERT INTO p2pdb.test6 (name2,email2,id) VALUES('123','123',123)"
	data.DatabaseName = "test6"
	data.TableName = "test6"
	data.DMLType = "insert"
	InsertStoreLog(*data)
}
