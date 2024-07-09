package model

import (
	"time"
)

const TableNameConfig = "config"

// Config 全局配置表
type Config struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Group_    string    `gorm:"column:group;type:varchar(32);not null;comment:配置分组" json:"group"` // 配置分组
	Key       string    `gorm:"column:key;type:varchar(32);not null;comment:配置的key" json:"key"`   // 配置的key
	Value     string    `gorm:"column:value;not null;comment:配置value" json:"value"`               // 配置value
	Intro     string    `gorm:"column:intro;type:varchar(64);not null;comment:配置说明" json:"intro"` // 配置说明
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*Config) TableName() string {
	return TableNameConfig
}
