package model

import (
	"time"
)

const TableNameArticle = "article"

// Article 文章表
type Article struct {
	Uuid       string    `gorm:"column:uuid;type:varchar(32);primaryKey;" json:"uuid"`
	CategoryId int32     `gorm:"column:category_id;not null;comment:分类" json:"category_id"`       //分类
	Title      string    `gorm:"column:title;type:varchar(256);not null;comment:标题" json:"title"` // 标题
	Cover      string    `gorm:"column:cover;type:varchar(1024);comment:封面" json:"cover"`         // 封面
	State      int32     `gorm:"column:state;not null;comment:审核状态" json:"state"`                 // 审核状态
	CommentNum int32     `gorm:"column:comment_num;not null;comment:浏览数" json:"comment_num"`      //评论数
	ViewNum    int32     `gorm:"column:view_num;not null;comment:浏览数" json:"view_num"`            // 浏览数
	LikeNum    int32     `gorm:"column:like_num;not null;comment:点赞数" json:"like_num"`            // 点赞数
	UserId     int32     `gorm:"column:user_id;not null;index;comment:用户ID" json:"user_id"`       //用户uuid
	IsHid      int32     `gorm:"column:is_hid;not null;comment:是否禁用：1是 0否" json:"is_hid"`         // 是否禁用：1是 0否
	IsDel      int32     `gorm:"column:is_del;not null;comment:是否删除：1是 0否" json:"is_del"`         // 是否删除：1是 0否
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  int64     `gorm:"column:deleted_at;not null;comment:删除时间" json:"deleted_at"` // 删除时间

	CategoryInfo ArticleCategory `gorm:"foreignKey:id;references:category_id" json:"category_info"`  // 分类详情
	DetailInfo   ArticleDetail   `gorm:"foreignKey:article_uuid;references:uuid" json:"detail_info"` // 文章详情
	UserInfo     User            `gorm:"foreignKey:id;references:user_id" json:"user_info"`          // 用户详情
}

func (*Article) TableName() string {
	return TableNameArticle
}
