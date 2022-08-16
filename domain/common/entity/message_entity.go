package entity

//module name+topic name+event
type Message struct {
	Type string
	Data Data
}

type Data struct {
	TableName     string
	DMLType       string
	DDLType       string
	DDLActionType string
	SQLStatement  string
}
