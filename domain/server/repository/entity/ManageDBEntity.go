package store

import "time"

// gorm.Model definition
type BaseColumn struct {
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// Node model definition
type Database struct {
	BaseColumn
	DatabaseId string `gorm:"column:database_id"`
}
