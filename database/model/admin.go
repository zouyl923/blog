package model

import (
	"time"
)

const TableNameAdmin = "admin"

// Admin 管理员表
type Admin struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(32);not null;comment:账号名" json:"name"`               // 账号名
	Phone     string    `gorm:"column:phone;type:char(11);not null;index:,unique;comment:手机号;" json:"phone"` // 手机号
	Password  string    `gorm:"column:password;type:varchar(64);not null;comment:密码" json:"password"`        // 密码
	RoleID    int32     `gorm:"column:role_id;type:int;not null;comment:角色ID" json:"role_id"`                // 角色ID
	IsHid     int32     `gorm:"column:is_hid;type:tinyint(1);not null;comment:是否禁用：1是 0否" json:"is_hid"`     // 是否禁用：1是 0否
	IsDel     int32     `gorm:"column:is_del;type:tinyint(1);not null;comment:是否删除：1是 0否" json:"is_del"`     // 是否删除：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间

	RoleInfo AdminRole `gorm:"foreignKey:id;references:role_id" json:"role_info"` // 角色信息
}

// TableName Admin table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
