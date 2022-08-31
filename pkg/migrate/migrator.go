package migrate

import (
	"gohub/pkg/database"
	"gorm.io/gorm"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
// 我们每个迁移都会给 migrations 表冲填充一个迁移记录
type Migration struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"varchar(255);not null;unique;"`
	Batch     int    // 记录迁移次数，迁移回滚时会用到
}

// NewMigrator 创建 Migrator 实例， 用以执行迁移操作
func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations", // 迁移文件存放目录
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
		// todo 为什么不直接使用 gorm.Migrator
	}

	// 创建 migrations 表，如果存在则无操作
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}

	return migrator
}

//func (migrator *Migrator) createMigrationsTable() {
//	migration := Migration{}
//
//	// 不存在则创建
//	if !migrator.Migrator.HasTable(&migration) {
//		migrator.Migrator.CreateTable(&migration)
//	}
//}
