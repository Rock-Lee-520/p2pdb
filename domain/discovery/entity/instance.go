package entity

type InstanceEntity struct {
	localPeerId     string
	localHost       string
	globalClockTime int64
}

func NewInstanceEntity() *InstanceEntity {
	return &InstanceEntity{}
}

func (i *InstanceEntity) GetLocalPeerId() string {
	return i.localPeerId
}

func (i *InstanceEntity) SetLocalPeerId(localPeerId string) {
	i.localPeerId = localPeerId
}

func (i *InstanceEntity) GetLocalHost() string {
	return i.localHost
}

func (i *InstanceEntity) SetLocalHost(localHost string) {
	i.localHost = localHost
}

func (i *InstanceEntity) GetGlobalClockTime() int64 {
	return i.globalClockTime
}

func (i *InstanceEntity) SetGlobalClockTime(globalClockTime int64) {
	i.globalClockTime = globalClockTime
}
