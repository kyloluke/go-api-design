package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Make api controller, example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	array := strings.Split(args[0], "/")
	// 处理参数，要求附带 API 版本（v1 或者 v2）
	if len(array) != 2 {
		console.Exit("err:api controller format: v1/user")
	}
	// 用 name 生成 model实例
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)
	// 组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/%s/%s_controller.go", apiVersion, model.TableName)
	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model)
}
