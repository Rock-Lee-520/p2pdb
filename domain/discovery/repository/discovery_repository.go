package persistence

import (
	PO "github.com/Rock-liyi/p2pdb/domain/discovery/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
)

var DB *orm.CreateDBFactory

func CreateInstance(name string) {

}

func GetInstance(name string) {

}

func CreateInstanceInformationTable() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.InstanceInformation{})
	return true
}

func CreatePeerNodeInfomation() bool {
	db := DB.InitDB()
	db.Migrator().CreateTable(&PO.PeerNodeInfomation{})
	return true
}
