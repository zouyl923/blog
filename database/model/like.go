package model

import (
	"time"
)

const TableNameLike = "like"

// Like 文章点赞表
type Like struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BizType   string    `gorm:"column:biz_type;not null;type:varchar(32);index;comment:业务类型" json:"biz_type"` //业务ID
	BizID     string    `gorm:"column:biz_id;not null;type:varchar(32);index;comment:业务ID" json:"biz_id"`     //关联ID
	UserId    int32     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                          //用户id
	IsLike    int32     `gorm:"column:is_like;not null;comment:是否喜欢：1是 0否" json:"is_like"`                    // 是否喜欢：1是 0否
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*Like) TableName() string {
	return TableNameLike
}
