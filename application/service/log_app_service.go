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
	logService.InitLogTable(data.TableName, data.TableName)
}

func InsertStoreLog(data commonEntity.Data, operation string) {
	log.Info("call InsertStoreLog method start")
	var nodeEntity = assembleNodeEntity()
	var dataEntity = assembleDataEntity(nodeEntity, operation, data)
	var linkEntity = assembleLinkEntity(nodeEntity)
	var dbEntity = assembleDbEntity(data)

	logService.InsertNodeLog(nodeEntity, dbEntity)
	logService.InsertLinkLog(linkEntity, dbEntity)
	logService.InsertDataLog(dataEntity, dbEntity)
	log.Info("call InsertStoreLog method end")
}

func assembleDbEntity(data commonEntity.Data) *entity.DB {
	var dbEntity = entity.NewDB()
	dbEntity.SetDatabaseName(data.GetDatabaseName())
	dbEntity.SetTableName(data.GetTableName())
	return dbEntity
}

func assembleNodeEntity() *entity.Node {
	var nodeEntity = entity.NewNode()
	nodeEntity.SetNodeId(function.GetUUID())
	nodeEntity.SetRequestId(function.GetUUID())
	nodeEntity.SetLogicalClock(function.GetGlobalLogicalClock() + 1)

	return nodeEntity
}

func assembleDataEntity(nodeEntity *entity.Node, operation string, d commonEntity.Data) *entity.Data {

	var dataEntity = entity.NewData()
	dataEntity.SetDataId(function.GetUUID())
	dataEntity.SetNodeId(nodeEntity.GetNodeId())
	dataEntity.SetOperation(operation)
	var Property entity.Properties
	var Properties []entity.Properties
	Property.SetSQLStatement(d.GetSQLStatement())
	Property.SetRowID(d.GetRowID())
	StringArray := append(Properties, Property)
	dataEntity.SetProperties(StringArray)
	return dataEntity
}

func assembleLinkEntity(nodeEntity *entity.Node) *entity.Link {
	var linkEntity = entity.NewLink()
	var instanceEntity = discoveryService.GetInstanceEntity()
	linkEntity.SetNodeId(nodeEntity.GetNodeId())
	linkEntity.SetLinkId(function.GetUUID())
	linkEntity.SetInstanceId(instanceEntity.GetInstanceId())
	// TODO: did not set real  values
	linkEntity.SetTableId(function.GetUUID())
	linkEntity.SetLastNodeId(function.GetUUID())
	linkEntity.SetDatabaseId(function.GetUUID())
	return linkEntity

}
