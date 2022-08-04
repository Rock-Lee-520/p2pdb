package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func InitInstanceInformation() {
	repository.CreateInstanceInformationTable()
	var InstanceId = function.GetLocalInstanceId()
	var LocalPeerId = function.GetLocalPeerId()
	var GlobalClockTime = function.GetGlobalLogicalClock()
	repository.InstanceInformationTable(InstanceId, LocalPeerId, GlobalClockTime)
}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomationTable()
}
