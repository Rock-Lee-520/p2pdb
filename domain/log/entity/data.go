package entity

type DataEntity struct {
	DataId     string   `json:"dataId"`
	Operation  string   `json:"operation"` //create  update delete
	Properties []string `json:"properties"`
}

func NewDataEntity() *DataEntity {
	return &DataEntity{}
}

func (i *DataEntity) GetDataId() string {
	return i.DataId
}

func (i *DataEntity) SetDataId(dataId string) {
	i.DataId = dataId
}

func (i *DataEntity) GetOperation() string {
	return i.Operation
}

func (i *DataEntity) SetOperation(operation string) {
	i.Operation = operation
}

func (i *DataEntity) GetProperties() string {
	return i.Operation
}

func (i *DataEntity) SetProperties(properties []string) {
	i.Properties = properties
}
