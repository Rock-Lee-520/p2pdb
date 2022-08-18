package po

// Link  information_schema definition
type Link struct {
	BaseColumn
	LinkId     string `gorm:"column:link_id"`
	LastNodeId string `gorm:"column:last_node_id"`
	NodeId     string `gorm:"column:node_id"`
	TableId    string `gorm:"column:table_id;primary_key"`
	DatabaseId string `gorm:"column:database_id"`
	InstanceId string `gorm:"column:local_instance_id"`
	TabName    string `gorm:"-"`
}

func NewLink() *Link {
	return &Link{}
}

func (n *Link) SetTabName(TabName string) {
	LINK_TABLE_NAME = TabName
}

func (n *Link) TableName() string {
	return LINK_TABLE_NAME
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
