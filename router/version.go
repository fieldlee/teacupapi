package router

import "teacupapi/controllers/system"

func initVersionRouter() {
	version := v1.Group("/version")
	version.POST("/getVersion", system.CheckUpdate)
}
