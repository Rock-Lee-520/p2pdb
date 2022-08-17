package entity

type LinkEntity struct {
	LinkId     string `json:"linkId"`
	LastNodeId string `json:"lastNodeId"`
	NodeId     string `json:"nodeId"`
	TableId    string `json:"tableId"`
	DatabaseId string `json:"databaseId"`
	InstanceId string `json:"instanceId"`
}

func NewLinkEntity() *LinkEntity {
	return &LinkEntity{}
}

func (i *LinkEntity) GetLinkId() string {
	return i.LinkId
}

func (i *LinkEntity) SetLinkId(linkId string) {
	i.LinkId = linkId
}

func (i *LinkEntity) GetNodeId() string {
	return i.NodeId
}

func (i *LinkEntity) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *LinkEntity) GetLastNodeId() string {
	return i.LastNodeId
}

func (i *LinkEntity) SetLastNodeId(lastNodeId string) {
	i.LastNodeId = lastNodeId
}

func (i *LinkEntity) GetTableId() string {
	return i.TableId
}

func (i *LinkEntity) SetTableId(tableId string) {
	i.TableId = tableId
}

func (i *LinkEntity) GetDatabaseId() string {
	return i.DatabaseId
}

func (i *LinkEntity) SetDatabaseId(databaseId string) {
	i.DatabaseId = databaseId
}

func (i *LinkEntity) GetInstanceId() string {
	return i.InstanceId
}

func (i *LinkEntity) SetInstanceId(instanceId string) {
	i.InstanceId = instanceId
}
