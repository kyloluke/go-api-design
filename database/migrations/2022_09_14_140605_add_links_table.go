package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	// 表名 默认取结构体的复数形式
	type Link struct {
		models.BaseModel
		Name string `gorm:"type:varchar(255);not null"`
		Link string `gorm:"type:varchar(255);not null"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_09_14_140605_add_links_table", up, down)
}
