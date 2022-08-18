package po

// Data  information_schema definition
type Data struct {
	BaseColumn
	DataId     string `gorm:"column:data_id"`
	NodeId     string `gorm:"column:node_id"`
	Operation  string `gorm:"column:operation"`
	Properties string `gorm:"column:properties"`
	TabName    string `gorm:"-"`
}

func NewData() *Data {
	return &Data{}
}

func (n *Data) SetTabName(TabName string) {
	//n.TabName = TabName
	DATA_TABLE_NAME = TabName
}

func (n *Data) TableName() string {
	return DATA_TABLE_NAME
}

func (i *Data) GetNodeId() string {
	return i.NodeId
}

func (i *Data) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *Data) GetDataId() string {
	return i.DataId
}

func (i *Data) SetDataId(dataId string) {
	i.DataId = dataId
}

func (i *Data) GetOperation() string {
	return i.Operation
}

func (i *Data) SetOperation(operation string) {
	i.Operation = operation
}

func (i *Data) GetProperties() string {
	return i.Operation
}

func (i *Data) SetProperties(properties string) {
	i.Properties = properties
}
