package entity

type Link struct {
	LinkId     string `json:"linkId"`
	LastNodeId string `json:"lastNodeId"`
	NodeId     string `json:"nodeId"`
	TableId    string `json:"tableId"`
	DatabaseId string `json:"databaseId"`
	InstanceId string `json:"instanceId"`
}

func NewLink() *Link {
	return &Link{}
}

func (i *Link) GetLinkId() string {
	return i.LinkId
}

func (i *Link) SetLinkId(linkId string) {
	i.LinkId = linkId
}

func (i *Link) GetNodeId() string {
	return i.NodeId
}

func (i *Link) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *Link) GetLastNodeId() string {
	return i.LastNodeId
}

func (i *Link) SetLastNodeId(lastNodeId string) {
	i.LastNodeId = lastNodeId
}

func (i *Link) GetTableId() string {
	return i.TableId
}

func (i *Link) SetTableId(tableId string) {
	i.TableId = tableId
}

func (i *Link) GetDatabaseId() string {
	return i.DatabaseId
}

func (i *Link) SetDatabaseId(databaseId string) {
	i.DatabaseId = databaseId
}

func (i *Link) GetInstanceId() string {
	return i.InstanceId
}

func (i *Link) SetInstanceId(instanceId string) {
	i.InstanceId = instanceId
}
