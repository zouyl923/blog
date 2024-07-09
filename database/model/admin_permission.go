package model

import (
	"time"
)

const TableNameAdminPermission = "admin_permission"

// AdminPermission 管理后台-菜单-关联权限表
type AdminPermission struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MenuID    int32     `gorm:"column:menu_id;not null;comment:菜单ID" json:"menu_id"`                               // 菜单ID
	URI       string    `gorm:"column:uri;type:varchar(256);not null;comment:权限访问-接口地址,可以多个可以换行，可以*匹配" json:"uri"` // 权限访问-接口地址,可以多个可以换行，可以*匹配
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminPermission's table name
func (*AdminPermission) TableName() string {
	return TableNameAdminPermission
}
