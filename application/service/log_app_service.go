package service

// DiscoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"

import (
	commonEntity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	discoveryService "github.com/Rock-liyi/p2pdb/domain/discovery/service"
	"github.com/Rock-liyi/p2pdb/domain/log/entity"
	logService "github.com/Rock-liyi/p2pdb/domain/log/service"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func InitLogTable(data commonEntity.Data) {
	// var instanceEntity = DiscoveryService.GetInstanceEntity()
	// log.Debug(instanceEntity)
	// var tableName="",
	// var DatabaseName=""
	log.Debug(data)
	logService.InitLogTable(data.TableName, data.TableName)
}

func InsertStoreLog(data commonEntity.Data) {
	var instanceEntity = discoveryService.GetInstanceEntity()

	var nodeEntity = entity.NewNodeEntity()
	var dataEntity = entity.NewDataEntity()
	var linkEntity = entity.NewLinkEntity()
	nodeEntity.SetNodeId(function.GetUUID())
	nodeEntity.SetRequestId(function.GetUUID())
	nodeEntity.SetLogicalClock(function.GetGlobalLogicalClock() + 1)
	logService.InsertNodeLog(nodeEntity)

	linkEntity.SetNodeId(nodeEntity.GetNodeId())
	linkEntity.SetLinkId(function.GetUUID())
	linkEntity.SetInstanceId(instanceEntity.InstanceId)
	// TODO: did not set real  values
	linkEntity.SetTableId(function.GetUUID())
	linkEntity.SetLastNodeId(function.GetUUID())
	linkEntity.SetDatabaseId(function.GetUUID())

	dataEntity.SetDataId(function.GetUUID())
	dataEntity.SetOperation(function.GetUUID())
	var StringArray []string
	StringArray = append(StringArray, function.GetUUID())
	dataEntity.SetProperties(StringArray)

	logService.InsertLinkLog(linkEntity)
	logService.InsertDataLog(dataEntity)
}
