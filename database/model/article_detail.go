package model

const TableNameArticleDetail = "article_detail"

// ArticleDetail 文章详情表
type ArticleDetail struct {
	ArticleUuid string `gorm:"column:article_uuid;type:varchar(32);primaryKey;" json:"article_uuid"`
	Content     string `gorm:"column:content;type:longtext;not null;comment:文章内容" json:"content"` // 文章内容
	Reason      string `gorm:"column:reason;type:varchar(256);comment:审核不通过原因" json:"reason"`     // 审核不通过原因
}

func (*ArticleDetail) TableName() string {
	return TableNameArticleDetail
}
