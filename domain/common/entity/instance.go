package entity

type InstanceEntity struct {
	localPeerId string
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
