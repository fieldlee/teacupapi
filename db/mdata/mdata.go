package mdata

import (
	"errors"
	"github.com/json-iterator/go"
)

//分页用
type PagingRep struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Count    int         `json:"totalRecord"`
	List     interface{} `json:"list"`
}

var (
	Cjson           = jsoniter.ConfigCompatibleWithStandardLibrary
	ErrNotFound     = errors.New("record not found")
	ErrTokenEmpty   = errors.New("token is empty")
	ErrMemberLocked = errors.New("member has been lock")
	ErrForbid       = errors.New("forbid")
	DomainConfigs   map[string]DomainConfig //api提供者 的域名配置
	PlatformPrefix  = ""                    //会员前缀
)

type DomainConfig struct {
	WebServer string `json:"webServer"`
}

// Query 查询字符串结构
type QueryString struct {
	Key   string
	Value string
}

// post附加数据
type AppendPostData struct {
	Key   string
	Value interface{}
}

const (
	AdminServer     = 1
	AgentServer     = 2
	ActiveSySConfig = 3
	FdAdmin         = 4
)

//http请求类型
const (
	HttpPost   = "POST"
	HttpDelete = "DELETE"
)

const (
	Android          = "android"
	IOS              = "ios"
	Web              = "web"
	RedisKeyNotExist = "redis: nil"
	HeaderPlatform   = "X-API-PLATFORM"
	HeaderToken      = "X-API-TOKEN"
	ClientType       = "Client-Type"
	Version          = "Version"
)

const (
	SpecialChar = `[ _~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]`
)
