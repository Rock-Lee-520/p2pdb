package service

// DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"

import (
	entity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	LogService "github.com/Rock-liyi/p2pdb/domain/log/service"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func InitLogTable(data entity.Data) {
	// var instanceEntity = DiscoveryService.GetInstanceEntity()
	// log.Debug(instanceEntity)
	// var tableName="",
	// var DatabaseName=""
	log.Debug(data)
	LogService.InitLogTable(data.TableName, data.TableName)
}

func InsertStoreLog() {

}
