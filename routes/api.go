package routes

import (
	v1controllers "gohub/app/http/controllers/v1"
	"gohub/app/http/controllers/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 注册一个路由
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("100-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("50-H")) // todo 为什么内层比外层的还多？？为什么这样设计呢
		{
			// 注册
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.LimitPerRoute("30-H"), middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.LimitPerRoute("30-H"), middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail) // 通过邮箱验证码注册

			// 发送验证码 图片验证码，需要加限流
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("30-H"), vcc.ShowCaptcha)  // 显示图片验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail) // 验证图片验证码，验证成功后，给邮箱发送验证码

			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			prc := new(auth.PasswordResetController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), prc.ResetByEmail)
		}
		uc := new(v1controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)

		// 用户
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
		}

		// 分类
		categoryController := new(v1controllers.CategoriesController)
		categoryGroup := v1.Group("/category")
		{
			categoryGroup.GET("", categoryController.Index)
			categoryGroup.POST("", categoryController.Store)
			categoryGroup.PUT("/:id", categoryController.Update)
			categoryGroup.DELETE("/:id", categoryController.Delete)
		}

		// 话题 帖子
		topicController := new(v1controllers.TopicsController)
		topicGroup := v1.Group("/topic")
		{
			topicGroup.POST("", middlewares.AuthJWT(), topicController.Store)
		}

	}
}
