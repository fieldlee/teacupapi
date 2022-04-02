package router

import "teacupapi/controllers/common"

// 初始化用户的一些路由
func initCommonRouter() {
	routerGroup := v1.Group("/common")
	{
		routerGroup.POST("/upload", common.FileUpload) //上传图片
	}
}
