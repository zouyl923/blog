package model

import (
	"time"
)

const TableNameAdminMenu = "admin_menu"

// AdminMenu 管理后台-菜单表
type AdminMenu struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int32     `gorm:"column:parent_id;not null;comment:父级ID" json:"parent_id"`         // 父级ID
	Name      string    `gorm:"column:name;type:varchar(32);not null;comment:菜单名称" json:"name"`  // 菜单名称
	URI       string    `gorm:"column:uri;type:varchar(32);comment:菜单地址" json:"uri"`             // 菜单地址
	Icon      string    `gorm:"column:icon;type:varchar(32);comment:iconfont图标库的图标" json:"icon"` // iconfont图标库的图标
	Weight    int32     `gorm:"column:weight;not null;comment:权重" json:"weight"`                 // 权重
	IsHid     int32     `gorm:"column:is_hid;not null;comment:是否禁用：1是 0否" json:"is_hid"`         // 是否禁用：1是 0否
	IsDel     int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"`         // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间

	Permission AdminPermission `gorm:"foreignKey:menu_id;references:id" json:"permission"` // 权限信息
}

// TableName AdminMenu's table name
func (*AdminMenu) TableName() string {
	return TableNameAdminMenu
}
