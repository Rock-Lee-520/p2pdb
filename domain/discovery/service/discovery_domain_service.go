package service

import (
	"time"

	"github.com/Rock-liyi/p2pdb/domain/discovery/entity"
	repository "github.com/Rock-liyi/p2pdb/domain/discovery/repository"
	function "github.com/Rock-liyi/p2pdb/infrastructure/util/function"
)

func GetInstanceEntity() *entity.InstanceEntity {
	var instance = repository.GetInstanceInformation()

	var instanceEntity = entity.NewInstanceEntity()
	instanceEntity.SetGlobalClockTime(instance.GlobalClockTime)
	instanceEntity.SetLocalHost(instance.LocalHost)
	instanceEntity.SetLocalPeerId(instance.LocalPeerId)
	instanceEntity.SetInstanceId(instance.InstanceId)
	return instanceEntity
}

func InitInstanceInformation(LocalPeerId string) string {
	repository.CreateInstanceInformationTable()

	var instance = repository.GetInstanceInformation()
	var instanceId string
	var GlobalClockTime = function.GetGlobalLogicalClock()
	var GetLocalHost = function.GetLocalHost()
	// if didn't find instance, create a new one
	if instance.InstanceId == "" {
		instanceId = function.GetLocalInstanceId()

		instance.InstanceId = instanceId
		instance.LocalPeerId = LocalPeerId
		instance.GlobalClockTime = GlobalClockTime
		instance.LocalHost = GetLocalHost
		instance.CreatedAt = time.Now()
		instance.UpdatedAt = time.Now()

		repository.InsertInstanceInformation(instance)

	} else {
		instance.LocalPeerId = LocalPeerId
		instance.GlobalClockTime = GlobalClockTime + 1
		instance.UpdatedAt = time.Now()
		//update instance information LocalPeerId and GlobalClockTime
		repository.UpdateInstanceInformation(instance)
	}

	return instanceId
}

func InitPeerNodeInfomation() {
	repository.CreatePeerNodeInfomationTable()
}
