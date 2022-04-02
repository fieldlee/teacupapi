// http 返回的状态码, 需要特别注意跳到登录窗口
package utils

const (
	Success = "success"
	Error   = "error"

	StatusOK             = 6000 // 正常返回
	ErrAccess            = 6001 // 登录状态失效 类似401, 会跳到登录窗口
	ErrAccount           = 6002 // 账号或密码错误
	ErrRefuse            = 6003 // 访问被拒绝类似403
	ErrNotFound          = 6004 // 找不到接口类似404
	ErrInternal          = 6005 // 接口服务器产生错误 类似500
	ErrInvalidGateway    = 6006 // 无效网关 类似502
	ErrAPIFailed         = 6007 // 接口请求失败, 无法接受参数，会跳到登录窗口
	ErrInvalidParams     = 6008 // 违法的参数
	ErrDataExistFailed   = 6009 // 数据已经存在
	ErrFileUploadFailed  = 6010 // 文件上传失败
	ErrLockedFailed      = 6011 // 频繁提交拒绝
	ErrWarning           = 6012 // 警告
	ErrMaintenance       = 6013 // 维护
	ErrCallError         = 6014 // 调用错误
	ErrNotBindPhone      = 6015 // 未绑定手机号的错误码
	ErrLoginVerify       = 6016 // 用户登录需要手机验证
	ErrToken             = 6018 // token解析错误
	ErrCodeError         = 6019 // 验证码错误
	ErrThirdError        = 6020 // 第三方验证错误
	ErrUserExistError    = 6021 // 用户已存在
	ErrPhoneError        = 6022 // 错误的手机号
	ErrDataError         = 6023 // 数据库存储失败
	ErrDataNotExistError = 6024 // 用户不存在
	ErrUserError         = 6025 // 用户权限不对，只有创建人才可以操作
	ErrCodeTimeError     = 6026 // 验证码操作太频繁
	ErrAllowError        = 6027 // 不允许操作

	InvalidJson   = "json解析失败"
	InvalidParams = "违法的参数"

	WrongPhoneNo  = "错误的手机号"
	ParamInvalid  = "参数错误"
	QueryFail     = "查询失败"
	UpdateFail    = "更新失败"
	AddFail       = "新增失败"
	DelFail       = "删除失败"
	UploadSuccess = "上传成功"
	UploadError   = "上传失败"
	QueryFailed   = "查询失败"
	QuerySuccess  = "查询成功"

	EmailExist       = "邮箱已经注册"
	NameExist        = "用户已经存在"
	InvalidKey       = "主键创建失败"
	EmptyName        = "用户名为空"
	EmptyUserId      = "用户ID为空"
	EmptyPwd         = "密码为空"
	EmptyGroup       = "组信息为空"
	InvalidName      = "用户名不符合规范"
	InvalidPwd       = "密码不符合规范"
	InvalidEmail     = "邮箱不符合规范"
	MerCodeExist     = "商户类型已经存在"
	EmptyRole        = "角色名为空"
	EmptyRoleId      = "角色ID为空"
	RoleExist        = "角色名已经存在"
	UserNotExist     = "用户不存在"
	RoleNotExist     = "角色不存在"
	AdminNoMerchent  = "管理员不能关联商户"
	GroupNoExist     = "不存在的组名称"
	EmptyRecord      = "record not found"
	EmptyRecord2     = "结果为空记录"
	AdminNoAdd       = "不能新增超级管理员"
	AdminNoEdit      = "不能修改超级管理员"
	MenuPageMulti    = "授权的角色权限重复"
	ListNoSuperAdmin = "不能查看超管信息"
	IpNoWhite        = "登录IP非白名单"
	EmailIsExist     = "邮箱已经存在"
	IpWhiteMulti     = "IP已经存在白名单中"
	PwdLessThanSix   = "密码长度应该大于等于6小于等于24"
	NameRegular      = "姓名长度2-20，由大小写字母以及数字组成"
	PwdRegular       = "密码长度6-24，由大小写字母以及数字组成,可包含_!-@,."
	InternalErr      = "内部错误"
	TokenErr         = "token解析错误"
)
