package model

import (
	"time"
)

const TableNameAdminMessage = "admin_message"

// AdminMessage 管理后台消息表
type AdminMessage struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdminID   int32     `gorm:"column:admin_id;not null;index;comment:接受者ID" json:"admin_id"`      // 接受者ID
	Title     string    `gorm:"column:title;type:varchar(256);not null;comment:消息标题" json:"title"` // 消息标题
	Content   string    `gorm:"column:content;not null;comment:消息内容" json:"content"`               // 消息内容
	IsRead    int32     `gorm:"column:is_read;not null;comment:是否已读：1是 0否" json:"is_read"`         // 是否已读：1是 0否
	IsDel     int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"`           // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间
}

// TableName AdminMessage's table name
func (*AdminMessage) TableName() string {
	return TableNameAdminMessage
}
