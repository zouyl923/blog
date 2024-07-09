package model

import (
	"time"
)

const TableNameLikeCount = "like_count"

// LikeCount 点赞统计表
type LikeCount struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BizType    string    `gorm:"column:biz_type;not null;type:varchar(32);index;comment:业务类型" json:"biz_type"` //业务ID
	BizID      string    `gorm:"column:biz_id;not null;type:varchar(32);index;comment:业务ID" json:"biz_id"`     //关联ID
	UserId     int32     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                          //用户id
	LikeNum    int32     `gorm:"column:like_num;not null;comment:喜欢数" json:"is_like"`                          // 喜欢数
	DisLikeNum int32     `gorm:"column:dislike_num;not null;comment:不喜欢数" json:"dislike_num"`                  // 不喜欢数
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*LikeCount) TableName() string {
	return TableNameLikeCount
}
