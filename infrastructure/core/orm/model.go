package store

import (
	"time"
)

// gorm.Model definition
type BaseColumn struct {
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

// Node model definition
type Node struct {
	BaseColumn
	NodeId             string `gorm:"column:node_id"`
	NodeType           string `gorm:"column:node_type"`
	LamportClock       int64  `gorm:"column:lamport_clock"`
	ReceivingTimestamp int32  `gorm:"column:receiving_timestamp"`
	ReceivingDate      string `gorm:"column:receiving_date"`
	SendingDate        string `gorm:"column:sending_date"`
	SendingTimestamp   int32  `gorm:"column:sending_timestamp"`
	LastNodeId         string `gorm:"column:last_node_id"`
}

//  Object model definition

type Object struct {
	BaseColumn
	ObjectId  string `gorm:"column:object_id"`
	NodeId    string `gorm:"column:node_id"`
	Content   string `gorm:"column:content"`
	Operation string `gorm:"column:operation"`
	Property  string `gorm:"column:propertie"`
}

//  Link model definition
type Link struct {
	BaseColumn
	LinkId     string `gorm:"column:link_id"`
	OldLinkId  string `gorm:"column:old_link_id"`
	LastNodeId string `gorm:"column:last_node_id"`
	NodeID     string `gorm:"column:node_id"`
	//LinkSize   string `gorm:"column:link_size"`
}

//  Link model definition
type Service struct {
	BaseColumn
	ServiceId string `gorm:"column:link_id"`
	Cid       string `gorm:"column:cid"`
	Ip        string `gorm:"column:ip"`
	//LinkSize   string `gorm:"column:link_size"`
}

// type User struct {
// 	gorm.Model
// 	Name         string
// 	Age          sql.NullInt64
// 	Birthday     *time.Time
// 	Email        string  `gorm:"type:varchar(100);unique_index"`
// 	Role         string  `gorm:"size:255"`        // 设置字段⼤⼩为255
// 	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯⼀并且不为空
// 	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为⾃增类型
// 	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
// 	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
// }
