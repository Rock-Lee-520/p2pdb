package service

import (
	DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"
	LogService "github.com/Rock-liyi/p2pdb/domain/log/service"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func CreateStoreLogTable() {
	var instanceEntity = DiscoveryService.GetInstanceEntity()
	log.Debug(instanceEntity)
	var tableName="",
	var DatabaseName=""
	LogService.CreateNodeTable(tableName,DatabaseName)
}

func InsertStoreLog() {

}

func CreateStoreLogTable() {

}
