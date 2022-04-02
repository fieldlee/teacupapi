package models

var (
	// 路由过滤
	RouterFilterMap = map[string]string{
		"/teacup/api/v1/health":             "1", // 健康检查
		"/teacup/api/v1/user/register":      "1", // 注册
		"/teacup/api/v1/user/login":         "1", // 登录
		"/teacup/api/v1/user/getSmsCode":    "1", // 用户短信验证码
		"/teacup/api/v1/user/checkSmsCode":  "1", // 用户短信验证码验证
		"/teacup/api/v1/user/checkUpdate":   "1", // 版本更新
		"/teacup/api/v1/version/getVersion": "1", // 版本更新
	}

	// 路由鉴权过滤
	RouterAuthWhiteMap = map[string]string{}
)

// ip 信息
type IPLoc struct {
	Country  string // 国家
	Province string // 省份
	City     string // 城市
}
