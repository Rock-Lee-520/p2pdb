package entity

type DB struct {
	DatabaseName string `json:"databaseName"`
	TableName    string `json:"tableName"`
}

func NewDB() *DB {
	return &DB{}
}

func (i *DB) GetDatabaseName() string {
	return i.DatabaseName
}

func (i *DB) SetDatabaseName(databaseName string) {
	i.DatabaseName = databaseName
}

func (i *DB) GetTableName() string {
	return i.TableName
}

func (i *DB) SetTableName(tableName string) {
	i.TableName = tableName
}
