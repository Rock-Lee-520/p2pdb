package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
)

func InitInstanceInformation() {
	repository.CreateInstanceInformationTable()
}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomation()
}
