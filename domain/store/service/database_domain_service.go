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

	var dbinfo = repository.FindDatabaseInformationTable(databaseId)
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

	var dbinfo = repository.FindDatabaseInformationTable(databaseId)

	//entity convert PO
	var databasePO = &PO.DatabaseInfomation{}
	databasePO.SetDatabaseId(databaseEntity.GetDatabaseId())
	databasePO.SetDatabaseName(databaseEntity.GetDatabaseName())
	databasePO.SetLocalInstanceId(databaseEntity.GetInstanceId())
	databasePO.SetLogicalClock(databaseEntity.GetLogicalClock())

	if dbinfo.GetDatabaseId() == "" {
		repository.InsertDatabaseInformation(databasePO)
	} else {
		repository.UpdateDatabaseInformation(databasePO)
	}

	return databaseId
}
