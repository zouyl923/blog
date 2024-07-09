package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户表
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nickname  string    `gorm:"column:nickname;type:varchar(32);not null;comment:昵称" json:"nickname"`       // 账号名
	Phone     int64     `gorm:"column:phone;type:char(11);not null;index,unique;comment:手机号;" json:"phone"` // 手机号
	Avatar    string    `gorm:"column:avatar;type:varchar(1024);not null;comment:头像;" json:"avatar"`        // 头像
	Password  string    `gorm:"column:password;type:varchar(64);not null;comment:密码" json:"password"`       // 密码
	IsHid     int32     `gorm:"column:is_hid;type:tinyint(1);not null;comment:是否禁用：1是 0否" json:"is_hid"`    // 是否禁用：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName  table name
func (*User) TableName() string {
	return TableNameUser
}
