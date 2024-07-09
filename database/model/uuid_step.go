package model

import (
	"time"
)

const TableNameUuidStep = "uuid_step"

// UuidStep uuid生成配置表
type UuidStep struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BizType   string    `gorm:"column:biz_type;type:varchar(32);not null;index:,unique;comment:业务类型" json:"biz_type"` // 业务类型
	Step      int32     `gorm:"column:step;type:int;not null;comment:每次生成多少;" json:"step"`                            // 每次生成多少
	Version   int64     `gorm:"column:version;type:int;not null;comment:版本号;" json:"version"`                         //版本号
	MaxId     int64     `gorm:"column:max_id;type:int;not null;comment:上次生成最大ID;" json:"max_id"`                      // 当前最大ID
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"update_at"`
}

func (*UuidStep) TableName() string {
	return TableNameUuidStep
}
