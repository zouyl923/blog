package model

import (
	"time"
)

const TableNameUuid = "uuid"

// Uuid uuid生成表
type Uuid struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BizType   string    `gorm:"column:biz_type;type:varchar(32);index:idx_id,unique;not null;comment:业务类型" json:"biz_type"` // 业务类型
	BizID     string    `gorm:"column:biz_id;type:varchar(32);not null;index:idx_id,unique;comment:业务ID;" json:"biz_id"`    // 业务ID
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (*Uuid) TableName() string {
	//按照每天生成uuid
	return TableNameUuid + time.Now().Format("20060102")
}
