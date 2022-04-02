// 存放一些公用的结构体
package models

import (
	"errors"

	"github.com/json-iterator/go"
)

var (
	Cjson = jsoniter.ConfigCompatibleWithStandardLibrary

	ErrNotFound   = errors.New("record not found")
	TokenEmptyErr = errors.New("token is empty")
	MemberLocked  = errors.New("member has been lock")
)

const (
	Android = "android"
	IOS     = "ios"

	RedisKeyNotExist = "redis: nil"

	HeaderToken = "X-API-TOKEN"
	ClientType  = "Client-Type"
	Version     = "Version"
)

const (
	SpecialChar = `[ _~!@#$%^&*()+=|{}':;',\\[\\].<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]`
)

type PagingRep struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Count    int         `json:"totalRecord"`
	List     interface{} `json:"list"`
}
