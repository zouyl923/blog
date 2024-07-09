package model

import (
	"time"
)

const TableNameAdminPassword = "admin_password"

// AdminPassword 管理员-密码修改记录
type AdminPassword struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdminID     int32     `gorm:"column:admin_id;not null;index;comment:接受者ID" json:"admin_id"`                  // 接受者ID
	OldPassword string    `gorm:"column:old_password;type:varchar(64);not null;comment:旧密码" json:"old_password"` // 旧密码
	NewPassword string    `gorm:"column:new_password;type:varchar(64);not null;comment:新密码" json:"new_password"` // 新密码
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminPassword's table name
func (*AdminPassword) TableName() string {
	return TableNameAdminPassword
}
