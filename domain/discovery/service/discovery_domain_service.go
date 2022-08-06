package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func InitInstanceInformation(LocalPeerId string) {
	repository.CreateInstanceInformationTable()

	var instance = repository.GetInstanceInformation()

	// if didn't find instance, create a new one
	//log.Debug(instance.InstanceId)
	if instance.InstanceId == "" {
		var InstanceId = function.GetLocalInstanceId()
		var GlobalClockTime = function.GetGlobalLogicalClock()
		repository.InsertInstanceInformation(InstanceId, LocalPeerId, GlobalClockTime)

	} else {
		//update instance information LocalPeerId and GlobalClockTime
		repository.UpdateInstanceInformation(instance.InstanceId, LocalPeerId, instance.GlobalClockTime)
	}

}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomationTable()
}
