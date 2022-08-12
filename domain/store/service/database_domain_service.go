package service

import (
	entity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func InitDatabaseInformation() {

	repository.CreateDatabaseInformationTable()
}

func SaveDatabaseInformation(databaseName string, instanceId string) string {
	log.Debug("=====SaveDatabaseInformation")

	var databaseId = entity.GetNewDatabaseId(databaseName)
	var dbinfo = repository.FindDatabaseInformationTable(databaseId)
	log.Debug(dbinfo)
	if dbinfo.DatabaseId == "" {
		repository.InsertDatabaseInformation(databaseId, databaseName, instanceId)
	} else {
		repository.UpdateDatabaseInformation(databaseId, databaseName, instanceId)
	}

	return databaseId
}
