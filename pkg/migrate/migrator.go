package migrate

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
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
	migrator.createMigrationsTable()

	return migrator
}

// Up 执行所有未迁移的文件
func (migrator *Migrator) Up() {
	// 读取所有的有效的，格式正确的迁移文件，并确保时间顺序
	migrateFiles := migrator.readAllMigrationFiles()

	// 获取当前批次的值
	batch := migrator.getBatch()

	// 获取所有的成功迁移的数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	runed := false
	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrateFiles {

		// 对比文件名称，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback() {
	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	// 回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 回退迁移，按照倒序执行迁移的 down 方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 标记是否真的有执行了迁移回退的操作
	runed := false
	for _, _migration := range migrations {
		// 友好提示
		console.Warning("rollback " + _migration.Migration)

		// 执行迁移文件的 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true
		// 回退成功了就删除掉这条记录
		migrator.DB.Delete(&_migration)

		// 打印运行状态
		console.Success("finish " + mfile.FileName)
	}
	return runed
}

// 执行迁移，执行迁移文件的 up 方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)
		// 执行 up 方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}

	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

// getBatch 获取当前这个批次的值, 从 migrations 表中获取最后那条记录的 batch 字段的值
func (migrator *Migrator) getBatch() int {
	// 默认为 1
	batch := 1

	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// 如果有值就 +1
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migration 目录下的所有文件，默认是文件名称排序
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// 除掉末尾的 .go
		fileName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))

		// 通过迁移文件的名称获取『MigrationFile』对象，
		mfile := getMigrationFile(fileName) // mfile 是 migration_file.go 中的 MigrationFile 结构体对象
		// 加个判断，确保迁移文件可用，再放进 migrateFiles 数组中
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	// 返回排序好的『MigrationFile』数组
	return migrateFiles
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}

	// 不存在则创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

// Reset 回滚所有迁移
func (migrator *Migrator) Reset() {

	var migrations = []Migration{}
	// 按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)

	// 回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

// Refresh 回滚所有迁移，并重新执行所有迁移
func (migrator *Migrator) Refresh() {
	migrator.Reset()
	migrator.Up()
}

// Fresh Drop 所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {
	// 获取数据库名称，用以提示
	dbname := database.CurrentDatabase()

	// 删除所有表
	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Success("clearup database " + dbname)

	// 重新创建 migrates 表
	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	migrator.Up()
}
