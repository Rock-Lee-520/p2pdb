package po

import "time"

/*
* TODO maybe it will caused variable  value mixed with others
 */
var NODE_TABLE_NAME string = "node"
var DATA_TABLE_NAME string = "data"
var LINK_TABLE_NAME string = "link"

// gorm.Model definition
type BaseColumn struct {
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
