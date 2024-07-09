package model

import (
	"time"
)

const TableNameAdminRole = "admin_role"

// AdminRole 管理员-角色表
type AdminRole struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(32);not null;comment:角色名称" json:"name"` // 角色名称
	IsHid     int32     `gorm:"column:is_hid;not null;comment:是否禁用：1是 0否" json:"is_hid"`        // 是否禁用：1是 0否
	IsDel     int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"`        // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间
}

// TableName AdminRole's table name
func (*AdminRole) TableName() string {
	return TableNameAdminRole
}
