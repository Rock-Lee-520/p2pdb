package persistence

// Table  information_schema  definition
type TableInfomation struct {
	BaseColumn
	TableName string `gorm:"column:table_name"`
	ClockTime int64  `gorm:"column:peer_id"`
	Version   int64  `gorm:"column:version"`
	Schema    string `gorm:"column:schema"` // Create the sql statement of the table

}
