package service

import (
	StoreService "github.com/Rock-liyi/p2pdb/domain/store/service"
)

func InitStore() {

	StoreService.IninDatabaseInformation()
	StoreService.IninTableInformation()
}
