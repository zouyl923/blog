package model

import (
	"time"
)

const TableNameAdminRolePermission = "admin_role_permission"

// AdminRolePermission 管理员-角色-关联权限表
type AdminRolePermission struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID    int32     `gorm:"column:role_id;not null;comment:角色ID" json:"role_id"` // 角色ID
	MenuID    int32     `gorm:"column:menu_id;not null;comment:菜单ID" json:"menu_id"` // 菜单ID
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminRolePermission's table name
func (*AdminRolePermission) TableName() string {
	return TableNameAdminRolePermission
}
