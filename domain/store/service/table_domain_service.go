package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
)

func InitTableInformation() {
	repository.CreateTableInformationTable()
}
