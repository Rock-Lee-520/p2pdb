package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/store/repository"
)

func IninDatabaseInformation() {
	repository.CreateDatabaseInformationTable()
}
