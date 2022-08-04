package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func InitInstanceInformation() {
	repository.CreateInstanceInformationTable()

	var instance = repository.GetInstanceInformation()

	// if didn't find instance, create a new one
	//log.Debug(instance.InstanceId)
	if instance.InstanceId == "" {
		var InstanceId = function.GetLocalInstanceId()
		var LocalPeerId = function.GetLocalPeerId()
		var GlobalClockTime = function.GetGlobalLogicalClock()
		repository.InstanceInformationTable(InstanceId, LocalPeerId, GlobalClockTime)

	} else {
		log.Debug(instance.InstanceId)
	}

}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomationTable()
}
