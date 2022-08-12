package entity

type InstanceEntity struct {
	LocalPeerId     string
	LocalHost       string
	GlobalClockTime int64
}

func NewInstanceEntity() *InstanceEntity {
	return &InstanceEntity{}
}

func (i *InstanceEntity) GetLocalPeerId() string {
	return i.LocalPeerId
}

func (i *InstanceEntity) SetLocalPeerId(localPeerId string) {
	i.LocalPeerId = localPeerId
}

func (i *InstanceEntity) GetLocalHost() string {
	return i.LocalHost
}

func (i *InstanceEntity) SetLocalHost(localHost string) {
	i.LocalHost = localHost
}

func (i *InstanceEntity) GetGlobalClockTime() int64 {
	return i.GlobalClockTime
}

func (i *InstanceEntity) SetGlobalClockTime(globalClockTime int64) {
	i.GlobalClockTime = globalClockTime
}
