package po

// Table  information_schema definition
type TableInfomation struct {
	BaseColumn
	TableId         string `gorm:"column:table_id;primary_key"`
	DatabaseId      string `gorm:"column:database_id"`
	TableName       string `gorm:"column:table_name"`
	LocalInstanceId string `gorm:"column:local_instance_id"`
	LogicalClock    int64  `gorm:"column:logical_clock"`
	Version         int64  `gorm:"column:version"`
	Schema          string `gorm:"column:schema"`
}

func (instance *TableInfomation) GetTablePrimary() string {
	return "table_id"
}

func (table *TableInfomation) GetTableId() string {
	return table.TableId
}

func (table *TableInfomation) SetTableId(tableId string) {
	table.TableId = tableId
}

func (table *TableInfomation) GetTableName() string {
	return table.TableName
}

func (table *TableInfomation) SetTableName(tableName string) {
	table.TableName = tableName
}

func (table *TableInfomation) GetDatabaseId() string {
	return table.DatabaseId
}

func (table *TableInfomation) SetDatabaseId(databaseId string) {
	table.DatabaseId = databaseId
}

func (table *TableInfomation) SetLocalInstanceId(localInstanceId string) {
	table.LocalInstanceId = localInstanceId
}

func (table *TableInfomation) GetLocalInstanceId() string {
	return table.LocalInstanceId
}

func (table *TableInfomation) SetLogicalClock(logicalClock int64) {
	table.LogicalClock = logicalClock
}

func (table *TableInfomation) GetLogicalClock() int64 {
	return table.LogicalClock
}

func (table *TableInfomation) SetVersion(version int64) {
	table.Version = version
}

func (table *TableInfomation) GetVersion() int64 {
	return table.Version
}

func (table *TableInfomation) SetSchema(schema string) {
	table.Schema = schema
}

func (table *TableInfomation) GetSchema() string {
	return table.Schema
}
