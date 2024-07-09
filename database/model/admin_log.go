package model

import (
	"time"
)

const TableNameAdminLog = "admin_log"

// AdminLog 管理员操作日志表
type AdminLog struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdminID   int32     `gorm:"column:admin_id;type:int;not null;index;comment:操作者" json:"admin_id"` // 操作者
	Name      string    `gorm:"column:name;type:varchar(32);not null;comment:操作行为" json:"name"`      // 操作行为
	URI       string    `gorm:"column:uri;type:varchar(64);not null;comment:操作接口地址" json:"uri"`      // 操作接口地址
	IP        string    `gorm:"column:ip;type:varchar(32);not null;comment:ip" json:"ip"`            // ip
	Agent     string    `gorm:"column:agent;type:varchar(128);not null;comment:agent" json:"agent"`  // agent
	Data      string    `gorm:"column:data;comment:操作数据JSON" json:"data"`                            // 操作数据JSON
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AdminLog's table name
func (*AdminLog) TableName() string {
	return TableNameAdminLog
}
