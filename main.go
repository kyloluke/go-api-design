package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化Logger
	bootstrap.SetupLogger()

	gin.SetMode(gin.ReleaseMode)
	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("yvU6Kvgb62sRyNUYVegh", "142974"), "正确的答案")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("yvU6Kvgb62sRyNUYVegh", "000000"), "错误的答案")

	router := gin.New()

	//router.GET("/test_auth", middlewares.AuthJWT(), func(c *gin.Context) {
	//	userModel := auth.CurrentUser(c)
	//	response.Data(c, userModel)
	//})
	// 路由绑定 1.中间件 2.注册路由
	bootstrap.SetupRoute(router)
	err := router.Run(":3000")

	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
