package service

import (
	entity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
	PO "github.com/Rock-liyi/p2pdb/domain/store/repository/po"
)

func InitTableInformation() {
	repository.CreateTableInformationTable()
}

func SaveTableInformation(tableEntity entity.Table) string {
	var tableId = tableEntity.GetNewTableId(tableEntity.GetTableName())
	tableEntity.SetDatabaseId(tableId)

	var tableInfo = repository.FindTableInformation(tableId)

	//entity convert PO
	var tablePO = &PO.TableInfomation{}
	tablePO.SetDatabaseId(tableEntity.GetDatabaseId())
	tablePO.SetTableName(tableEntity.GetTableName())
	tablePO.SetTableId(tableEntity.GetTableId())
	tablePO.SetLocalInstanceId(tableEntity.GetLocalInstanceId())

	if tableInfo.GetDatabaseId() == "" {
		tablePO.SetLogicalClock(tableEntity.GetLogicalClock())
		tablePO.SetSchema(tableEntity.GetSchema())
		repository.InsertTableInformation(tablePO)
	} else {
		tablePO.SetLogicalClock(tableInfo.GetLogicalClock() + 1)
		tablePO.SetVersion(tableInfo.GetVersion() + 1)
		repository.UpdateTableInformation(tablePO)
	}

	return tableId
}
