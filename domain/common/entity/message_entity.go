package entity

//module name+topic name+event
type Message struct {
	Type string
	Data Data
}

type Data struct {
	TableName     string
	DatabaseName  string
	DMLType       string
	DDLType       string
	DDLActionType string
	SQLStatement  string
	RowID         string
}

func NewData() *Data {
	return &Data{}
}

func (i *Data) GetTableName() string {
	return i.TableName
}

func (i *Data) SetTableName(tableName string) {
	i.TableName = tableName
}

func (i *Data) GetDatabaseName() string {
	return i.DatabaseName
}

func (i *Data) SetDatabaseName(databaseName string) {
	i.DatabaseName = databaseName
}

func (i *Data) GetDMLType() string {
	return i.DMLType
}

func (i *Data) SetDMLType(DMLType string) {
	i.DMLType = DMLType
}

func (i *Data) GetDDLType() string {
	return i.DDLType
}

func (i *Data) SetDDLType(DDLType string) {
	i.DDLType = DDLType
}

func (i *Data) GetDDLActionType() string {
	return i.DDLActionType
}

func (i *Data) SetDDLActionType(DDLActionType string) {
	i.DDLActionType = DDLActionType
}

func (i *Data) GetSQLStatement() string {
	return i.SQLStatement
}

func (i *Data) SetSQLStatement(SQLStatement string) {
	i.SQLStatement = SQLStatement
}

func (i *Data) GetRowID() string {
	return i.RowID
}

func (i *Data) SetRowID(rowID string) {
	i.RowID = rowID
}
