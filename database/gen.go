package main

import (
	"blog/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var Dir = "../"

var DB *gorm.DB

func main() {
	NewDb()
	Migrate()
}

/**
 * 根据结构体生成表
 */
func Migrate() {

	DB.Migrator().AutoMigrate(
		&model.Admin{},
		&model.AdminLog{},
		&model.AdminMenu{},
		&model.AdminMessage{},
		&model.AdminPassword{},
		&model.AdminPermission{},
		&model.AdminRole{},
		&model.AdminRolePermission{},
		&model.Config{},

		&model.User{},

		&model.Article{},
		&model.ArticleDetail{},
		&model.ArticleCategory{},

		&model.Like{},
		&model.LikeCount{},

		&model.UuidStep{},
		&model.Uuid{},
	)
}

/**
 * 根据表生成model
 */
func Generator() {
	g := gen.NewGenerator(gen.Config{
		// 生成目录
		OutPath:      Dir + "query",
		ModelPkgPath: Dir + "model",
		// generate mode
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	// reuse your gorm db
	g.UseDB(DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	// 执行并生成代码
	g.Execute()
}

func NewDb() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=true&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//禁用生成外键关联
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
