package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CMDMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Make model's seeder file, example: make seeder",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeeder(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filePath := fmt.Sprintf("database/seeders/%s_seeder.go", model.PackageName)
	createFileFromStub(filePath, "seeder", model)
}
