package service

import (
	entity "github.com/Rock-liyi/p2pdb/domain/store/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
)

func InitTableInformation() {
	repository.CreateTableInformationTable()
}

func AddTableRecordInTableInformation(databaseEntity entity.Database) {
	repository.CreateTableInformationTable()
}
