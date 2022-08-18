package entity

type Data struct {
	DataId     string   `json:"dataId"`
	NodeId     string   `json:"nodeId"`
	Operation  string   `json:"operation"` //create  update delete
	Properties []string `json:"properties"`
}

func NewData() *Data {
	return &Data{}
}

func (i *Data) GetDataId() string {
	return i.DataId
}

func (i *Data) SetDataId(dataId string) {
	i.DataId = dataId
}

func (i *Data) GetNodeId() string {
	return i.NodeId
}

func (i *Data) SetNodeId(nodeId string) {
	i.NodeId = nodeId
}

func (i *Data) GetOperation() string {
	return i.Operation
}

func (i *Data) SetOperation(operation string) {
	i.Operation = operation
}

func (i *Data) GetProperties() []string {
	return i.Properties
}

func (i *Data) SetProperties(properties []string) {
	i.Properties = properties
}
