package cmd

import (
	"github.com/spf13/cobra"
	"gohub/database/migrations"
	"gohub/pkg/migrate"
)

var CMDMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	// 所有 migrate 下的子命令都会执行以下代码
}

var CMDMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run nmigrated migrations",
	Run:   runUp,
}

var CMDMigrateDown = &cobra.Command{
	Use:   "rollback",
	Short: "Reverse the up command",
	Run:   runDown,
}

var CMDMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollbacl all database migrations",
	Run:   runReset,
}

var CMDMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

var CMDMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all tables and re-run migrations",
	Run:   runFresh,
}

func init() {
	CMDMigrate.AddCommand(
		CMDMigrateUp,
		CMDMigrateDown,
		CMDMigrateReset,
		CMDMigrateRefresh,
		CMDMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

func runFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
