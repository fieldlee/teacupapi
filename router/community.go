package router

import (
	"teacupapi/controllers/community"
)

// 初始化用户的一些路由
func initCommunityRouter() {
	routerGroup := v1.Group("/community")
	{
		routerGroup.POST("/newCommunity", community.NewCommunity)           //创建社区
		routerGroup.POST("/updateCommunity", community.UpdateCommunity)     //更新社区
		routerGroup.POST("/getCommunity", community.GetCommunity)           //查找社区
		routerGroup.POST("/getHotCommunity", community.GetHotsCommunity)    //热门社区
		routerGroup.POST("/getVipsCommunity", community.GetVipsCommunity)   //我做嘉宾的社区
		routerGroup.POST("/getFansCommunity", community.GetFansCommunity)   //我关注的社区
		routerGroup.POST("/getMyCommunity", community.GetMyCommunity)       //我创建的社区
		routerGroup.POST("/getJoinCommunity", community.GetJoinCommunity)   //我加入的社区
		routerGroup.POST("/getCommunityFans", community.GetCommunityFans)   //社区的粉丝
		routerGroup.POST("/getCommunityJoins", community.GetCommunityJoins) //社区的成员
		routerGroup.POST("/getCommunityVips", community.GetCommunityVips)   //社区的嘉宾

		routerGroup.POST("/applyCommunity", community.Apply2Community)          //关注
		routerGroup.POST("/cancelCommunity", community.CloseCommunity)          //取消关注
		routerGroup.POST("/invite2Community", community.Invite2Community)       //邀请进入社区嘉宾
		routerGroup.GET("/getInvite2Community", community.GetInvite2Community)  //邀请进入社区嘉宾
		routerGroup.POST("/update2Community", community.Update2Community)       //邀请进入社区嘉宾
		routerGroup.POST("/joinCommunity", community.Join2Community)            //申请加入
		routerGroup.GET("/getJoinCommunity", community.GetJoin2Community)       //获得加入申请
		routerGroup.POST("/updateJoinCommunity", community.UpdateJoinCommunity) //审批申请加入
	}
}
