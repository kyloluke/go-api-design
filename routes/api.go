package routes

import (
	"gohub/app/http/controllers/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 注册一个路由
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			// 注册
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail) // 通过邮箱验证码注册

			// 发送验证码 图片验证码，需要加限流
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)  // 显示图片验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail) // 携带正确的图片验证码和输入的邮箱，给邮箱发送验证码

			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			// 重置密码
			prc := new(auth.PasswordResetController)
			authGroup.POST("/password-reset/using-email", prc.ResetByEmail)
		}

	}

}
