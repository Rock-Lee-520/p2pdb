package service

import (
	entity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
)

func InitDatabaseInformation() {

	repository.CreateDatabaseInformationTable()
}

func GetDatabaseInformation(databaseName string) *entity.Database {
	var databaseEntity = entity.NewDatabase()
	var databaseId = databaseEntity.GetNewDatabaseId(databaseName)
	databaseEntity.SetDatabaseId(databaseId)

	var dbinfo = repository.FindDatabaseInformation(databaseId)
	databaseEntity.SetDatabaseId(dbinfo.GetDatabaseId())
	databaseEntity.SetDatabaseName(dbinfo.GetDatabaseName())
	databaseEntity.SetInstanceId(dbinfo.GetLocalInstanceId())
	databaseEntity.SetLogicalClock(dbinfo.GetLogicalClock())
	return databaseEntity
}

func SaveDatabaseInformation(databaseName string, instanceId string) string {
	var databaseEntity = entity.NewDatabase()
	var databaseId = databaseEntity.GetNewDatabaseId(databaseName)
	databaseEntity.SetDatabaseId(databaseId)

	var dbinfo = repository.FindDatabaseInformation(databaseId)

	//entity convert PO
	var databasePO = &PO.DatabaseInfomation{}
	databasePO.SetDatabaseId(databaseEntity.GetDatabaseId())
	databasePO.SetDatabaseName(databaseEntity.GetDatabaseName())
	databasePO.SetLocalInstanceId(instanceId)

	if dbinfo.GetDatabaseId() == "" {
		databasePO.SetLogicalClock(databaseEntity.GetLogicalClock())
		repository.InsertDatabaseInformation(databasePO)
	} else {
		databasePO.SetLogicalClock(dbinfo.GetLogicalClock() + 1)
		repository.UpdateDatabaseInformation(databasePO)
	}

	return databaseId
}
