package persistence

import (
	PO "github.com/Rock-liyi/p2pdb/domain/discovery/repository/po"
	orm "github.com/Rock-liyi/p2pdb/infrastructure/core/orm"
	log "github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	//function "github.com/Rock-liyi/p2pdb/infrastructure/untl/function"
)

var ORM *orm.CreateDBFactory = new(orm.CreateDBFactory)
var db orm.DBconnect

func InitDB() orm.DBconnect {
	//settings  use  the default  information database
	ORM.SetIsInternalStore(true)
	db = ORM.InitDB()
	return db
}

func CreateInstance(name string) {

}

func GetInstance(name string) {

}

func CreateInstanceInformationTable() bool {
	db = InitDB()
	db.Migrator().CreateTable(&PO.InstanceInformation{})
	return true
}

func CreatePeerNodeInfomationTable() bool {
	db = InitDB()
	db.Migrator().CreateTable(&PO.PeerNodeInfomation{})
	return true
}

func InsertInstanceInformation(instance *PO.InstanceInformation) bool {
	db = InitDB()
	err := db.InsertIgnore(instance)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func UpdateInstanceInformation(instance *PO.InstanceInformation) bool {
	db = InitDB()

	err := db.Where(instance.GetInstancePrimaryId()+" = ?", instance.InstanceId).Updates(instance)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func GetInstanceInformation() *PO.InstanceInformation {
	db = InitDB()
	var instance = &PO.InstanceInformation{}
	db.First(instance)

	return instance
}
