package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

type User struct {
	models.BaseModel
}
type Category struct {
	models.BaseModel
}

func init() {

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;not null"`
		UserID     string `gorm:"type:bigint;not null;index"`
		CategoryID string `gorm:"type:bigint;not null;index"`

		// 表级别的外键约束，即使不声明，也不会影响 Preload(clause.Associations) 的调用
		User     User     // todo 这里为什么不能使用 user.User
		Category Category // todo 这里为什么不能使用 category.Category
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_09_12_234902_add_topic_table", up, down)
}
