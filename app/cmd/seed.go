package cmd

import (
	"github.com/spf13/cobra"
	"gohub/database/seeders"
	"gohub/pkg/console"
	"gohub/pkg/seed"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Inter fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // 最多传一个参数，可以不传
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 执行单个seeder
		name := args[0]
		sdr := seed.GetSeeder(name)
		if len(sdr.Name) > 0 {
			seed.RunSeeder(sdr)
		} else {
			console.Error("Seeder not found: " + name)
		}

	} else {
		// 执行所有的seeder
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
