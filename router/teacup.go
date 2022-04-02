package router

import "teacupapi/controllers/teacup"

func initTeacupRouter() {
	routerGroup := v1.Group("/teacup")
	{
		routerGroup.POST("/newTeacup", teacup.NewTeacup)            //创建茶室
		routerGroup.POST("/updateTeacup", teacup.UpdateTeacup)      //更新茶室
		routerGroup.POST("/myTeacup", teacup.MyTeacups)             //我的茶室
		routerGroup.POST("/getTeacup", teacup.GetTeacup)            //查找茶室
		routerGroup.POST("/invite2Teacup", teacup.Invite2Teacup)    //邀请嘉宾茶室
		routerGroup.GET("/getInviteTeacup", teacup.GetInviteTeacup) //查找邀请记录
		routerGroup.POST("/update2Teacup", teacup.Update2Teacup)    //邀请审批嘉宾茶室
		routerGroup.POST("/joinTeacup", teacup.Join2Teacup)         //加入茶室
		routerGroup.POST("/leaveTeacup", teacup.LeaveTeacup)        //离开茶室
		routerGroup.POST("/logTeacup", teacup.CreateTeacupHistory)  //开播记录茶室
		routerGroup.POST("/startTeacup", teacup.StartTeacup)        //开播
		routerGroup.POST("/closeTeacup", teacup.CloseTeacup)        //开播
		routerGroup.POST("/chatTeacup", teacup.CreateChat)          //开播
		routerGroup.POST("/markTeacup", teacup.CreateMark)          //打分
		routerGroup.POST("/teacupOfVips", teacup.VipTeacups)        //我作为嘉宾
		routerGroup.POST("/vipsOfTeacup", teacup.TeacupsVip)        //茶室的嘉宾
	}
}
