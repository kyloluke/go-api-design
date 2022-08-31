package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Make request file, example: make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1),
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/requests/%s_request.go", model.PackageName)

	createFileFromStub(filePath, "request", model)
}
