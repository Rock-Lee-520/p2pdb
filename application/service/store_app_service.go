package service

import (
	commonEntity "github.com/Rock-liyi/p2pdb/domain/common/entity"
	storeEntity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	StoreService "github.com/Rock-liyi/p2pdb/domain/store/service"
	conf "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func InitStore(instanceId string) {
	StoreService.InitTableInformation()

	StoreService.InitDatabaseInformation()
	var databaseName = conf.GetDBName()

	StoreService.SaveDatabaseInformation(databaseName, instanceId)
}

func AddTableRecord(data commonEntity.Data) {
	log.Debug("call AddTableRecord method start ")

	//databaseInformation=GetDatabaseInformation
	databaseEntity := StoreService.GetDatabaseInformation(data.GetDatabaseName())

	log.Debug(databaseEntity)

	tableEntity := storeEntity.NewTable()
	tableEntity.SetDatabaseId(databaseEntity.GetDatabaseId())
	tableEntity.SetTableId(tableEntity.GetNewTableId(data.GetTableName()))
	tableEntity.SetTableName(data.GetTableName())
	tableEntity.SetLocalInstanceId(databaseEntity.GetInstanceId())
	tableEntity.SetLogicalClock(databaseEntity.GetLogicalClock())
	StoreService.SaveTableInformation(*tableEntity)
	log.Debug("call AddTableRecord method end ")
}
