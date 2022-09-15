package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/app/cmd"
	"gohub/app/cmd/make"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()
		},
	}
	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CMDMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)
	// 配置默认运行 Web 服务  默认运行  server.go 注册的命令
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)
	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	// 配置初始化，依赖命令行 --env 参数
	//var env string
	//flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	//flag.Parse()
	//config.InitConfig(env)
	//
	//// 初始化Logger
	//bootstrap.SetupLogger()
	//
	//gin.SetMode(gin.ReleaseMode)
	//// 初始化 DB
	//bootstrap.SetupDB()

	//// 初始化 Redis
	//bootstrap.SetupRedis()

	//router := gin.New()

	//// 路由绑定 1.中间件 2.注册路由
	//bootstrap.SetupRoute(router)
	//err := router.Run(":3000")

	//if err != nil {
	//	// 错误处理，端口被占用了或者其他错误
	//	fmt.Println(err.Error())
	//}
}
