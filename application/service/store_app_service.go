package service

import (
	StoreService "github.com/Rock-liyi/p2pdb/domain/store/service"
	conf "github.com/Rock-liyi/p2pdb/infrastructure/util/config"
)

func InitStore(instanceId string) {
	StoreService.InitTableInformation()

	StoreService.InitDatabaseInformation()
	var databaseName = conf.GetDBName()

	StoreService.SaveDatabaseInformation(databaseName, instanceId)
}
