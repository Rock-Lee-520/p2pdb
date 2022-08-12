package service

import (
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func InitInstanceInformation(LocalPeerId string) string {
	repository.CreateInstanceInformationTable()

	var instance = repository.GetInstanceInformation()
	var instanceId string

	// if didn't find instance, create a new one
	if instance.InstanceId == "" {
		instanceId = function.GetLocalInstanceId()
		var GlobalClockTime = function.GetGlobalLogicalClock()
		repository.InsertInstanceInformation(instanceId, LocalPeerId, GlobalClockTime)

	} else {
		instanceId = instance.InstanceId
		//update instance information LocalPeerId and GlobalClockTime
		repository.UpdateInstanceInformation(instance.InstanceId, LocalPeerId, instance.GlobalClockTime)
	}

	return instanceId
}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomationTable()
}
