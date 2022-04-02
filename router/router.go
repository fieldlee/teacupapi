package router

import (
	"github.com/gin-gonic/gin"
	"teacupapi/controllers/base"
)

const (
	versionOne = "/teacup/api/v1"
	versionTwo = "/live/api/v2"
)

var (
	v1 *gin.RouterGroup
	v2 *gin.RouterGroup
)

// ApiRouter 路由
func ApiRouter(router *gin.Engine) {
	authorized := router.Group("/")
	authorized.Use(base.AuthRequired()) // 用户鉴权

	v1 = authorized.Group(versionOne)
	{
		v1.GET("/health", base.HealthCheck) // 健康检查
	}
	initCommonRouter()    // 上传图片
	initUserRouter()      //  用户路由
	initCommunityRouter() //社区路由
	initTeacupRouter()    //茶室路由
	initVersionRouter()   //版本路由

}
