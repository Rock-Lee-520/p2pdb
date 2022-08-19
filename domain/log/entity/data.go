package entity

type Data struct {
	DataId     string       `json:"dataId"`
	NodeId     string       `json:"nodeId"`
	Operation  string       `json:"operation"` //create  update delete
	Properties []Properties `json:"properties"`
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

func (i *Data) GetProperties() []Properties {
	return i.Properties
}

func (i *Data) SetProperties(properties []Properties) {
	i.Properties = properties
}

type Properties struct {
	SQLStatement string       `json:"sqlStatement"`
	FieldValue   []FieldValue `json:"fieldValue"`
	RowID        string       `json:"rowId"`
}

type FieldValue struct {
	Field string
	Value string
}

func (i *Properties) GetSQLStatement() string {
	return i.SQLStatement
}

func (i *Properties) SetSQLStatement(SQLStatement string) {
	i.SQLStatement = SQLStatement
}

func (i *Properties) GetFieldValue() []FieldValue {
	return i.FieldValue
}

func (i *Properties) SetFieldValue(fieldValue []FieldValue) {
	i.FieldValue = fieldValue
}

func (i *Properties) GetRowID() string {
	return i.RowID
}

func (i *Properties) SetRowID(rowID string) {
	i.RowID = rowID
}
