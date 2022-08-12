package service

import (
	entity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
)

func InitDatabaseInformation() {

	repository.CreateDatabaseInformationTable()
}

func SaveDatabaseInformation(databaseName string, instanceId string) string {
	var instanceEntity = entity.NewDatabaseEntity()
	var databaseId = entity.GetNewDatabaseId(databaseName)
	instanceEntity.SetDatabaseId(databaseId)

	var dbinfo = repository.FindDatabaseInformationTable(databaseId)

	//entity convert PO
	var databasePO = &PO.DatabaseInfomation{}
	databasePO.DatabaseId = instanceEntity.GetDatabaseId()
	databasePO.DatabaseName = instanceEntity.GetDatabaseName()
	databasePO.LocalInstanceId = instanceEntity.GetInstanceId()

	if dbinfo.DatabaseId == "" {

		repository.InsertDatabaseInformation(databasePO)
	} else {
		repository.UpdateDatabaseInformation(databasePO)
	}

	return databaseId
}
