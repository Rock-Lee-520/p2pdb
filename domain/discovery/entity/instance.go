package entity

type Instance struct {
	LocalPeerId string
	InstanceId  string
	LocalHost   string
	// GlobalClockTime int64
	LocalPublicKey  string
	LocalPrivateKey string
}

func NewInstance() *Instance {
	return &Instance{}
}

func (i *Instance) GetInstanceId() string {
	return i.InstanceId
}

func (i *Instance) SetInstanceId(instanceId string) {
	i.InstanceId = instanceId
}

func (i *Instance) GetLocalPeerId() string {
	return i.LocalPeerId
}

func (i *Instance) SetLocalPeerId(localPeerId string) {
	i.LocalPeerId = localPeerId
}

func (i *Instance) GetLocalHost() string {
	return i.LocalHost
}

func (i *Instance) SetLocalHost(localHost string) {
	i.LocalHost = localHost
}

// func (i *Instance) GetGlobalClockTime() int64 {
// 	return i.GlobalClockTime
// }

// func (i *Instance) SetGlobalClockTime(globalClockTime int64) {
// 	i.GlobalClockTime = globalClockTime
// }
