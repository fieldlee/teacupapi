package router

import (
	"teacupapi/controllers/user"
)

// 初始化用户的一些路由
func initUserRouter() {
	routerGroup := v1.Group("/user")
	{
		routerGroup.POST("/getSmsCode", user.GetSmsCode)           //用户短信验证码
		routerGroup.POST("/checkSmsCode", user.CheckSmsCode)       //短信验证码验证
		routerGroup.POST("/register", user.Register)               //注册
		routerGroup.POST("/login", user.Login)                     //登录
		routerGroup.POST("/loginOut", user.LoginOut)               //登出
		routerGroup.POST("/fillInformation", user.FillInformation) //完善资料
	}
}
