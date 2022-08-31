package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

// migrationFunc 定义 up 喝 down 回调方法的类型
type migrationFunc func(gorm.Migrator, *sql.DB)

// MigrationFile 代表单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// migrationFiles 所有的迁移文件的数组（切片）
var migrationFiles []MigrationFile

// Add 此方法新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
