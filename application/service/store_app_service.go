package service

import (
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

func CreateTableByStoreEvent(data interface{}) {
	log.Debug("call CreateTableByStoreEvent method start ")
	log.Debug(data)

	log.Debug("call CreateTableByStoreEvent method end ")
}
