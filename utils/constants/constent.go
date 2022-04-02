package constants

const (
	SpecialChar             = `[ _~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]`
	DefaultPage             = 1
	DefaultPageSize         = 15
	IpRegChar               = `^((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1 -9]?\d))))$`
	IpGroupRegChar          = `^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}-((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})$`
	REDIS_IP_WHITE2LIST     = "IP_WHITE2LIST"
	REDIS_IP_WHITELIST      = "IP_WHITELIST"
	REDIS_IP_BLACK          = "IP_BLACKELIST"
	REDIS_SINGLE_LOGIN      = "_TOKEN_LIST"
	REDIS_MULTI_LOGIN       = "MULTI_LOGIN_NAME"
	REDIS_WS_PULL_EXPIRE    = "WS_PULL_EXPIRE_"
	REDIS_WS_PUSH_EXPIRE    = "WS_PUSH_EXPIRE_"
	PageSize                = 10
	TotalMonthlyFeePreKey   = "ComCost_%s"                       //ComCost_topName
	TotalMonthlyFeeFieldKey = "%s_%d-%s"                         //topName_year-month
	NameLenMin              = 2                                  //用户名最小长度
	NameLenMax              = 20                                 //用户最大长度
	PwdLenMin               = 6                                  //密码最小长度
	PwdLenMax               = 12                                 //密码最大长度
	RegexpCheckName         = "^[a-zA-Z0-9]*$"                   //用户名校验规规则
	RegexpCheckPwd          = "^[a-zA-Z0-9\\_\\!\\-\\@\\,\\.]*$" //用户密码校验规则
	AppRegexpCheckPwd       = "^[A-Za-z0-9]{6,12}$"              //APP用户密码校验规则
	ChinaPhoneLen           = 11                                 //手机号长度
	USAPhoneLen             = 10                                 //手机号长度

	Checking         = 1 // 审批中
	Approved         = 2 // 审批通过
	Refused          = 3 // 拒绝
	SubjectForListen = "socket_group_subject"
)

const (
	REGION_CHINA = "cn"
	REGION_USA   = "us"
)
