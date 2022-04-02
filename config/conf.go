/*******
*统一配置文件中心
*******/

package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
	confInfo *Config
)

// Config info
type Config struct {
	BaseConf   `ini:"BaseConf"`
	RedisConf  `ini:"RedisConf"`
	RRedisConf `ini:"RRedisConf"`
	TybDbConf  `ini:"TybDbConf"`
	ReDbConf   `ini:"ReDbConf"`
	LogConf    `ini:"LogConf"`
	ConfigAddr `ini:"ConfigAddr"`
	UploadConf `init:"UploadConf"`
	NatsConfig `init:"NatsConfig"`
	APNSConf   `init:"APNSConfig"`
	S3Conf   `init:"S3Config"`
}

type NatsConfig struct {
	NatsUrl     string `ini:"NatsUrl"`
	ClusterID   string `ini:"ClusterID"`
	PubClientId string `ini:"PubClientId"`
	SubClientId string `ini:"SubClientId"`
}

type UploadConf struct {
	ImagesUrl        string `ini:"ImgsURL"`
	VideoURL         string `ini:"VideoURL"`
	FileStoragePath  string `ini:"fileStoragePath"`
	MaxUploadSize    int64  `ini:"maxUploadSize"`
	VideoStoragePath string `ini:"videoStoragePath"`
}

type BaseConf struct {
	Env             string `ini:"Env"`      // 环境信息
	HttpPort        string `ini:"HttpPort"` // http port
	AuthFlag        bool   `ini:"AuthFlag"` // 鉴权开关
	AesSert         string `ini:"AesSert"`  // aes 密钥
	SqlDB           string `ini:"SqlDB"`    // sql db
	SecretKey       string `ini:"SecretKey"`
	TokenKey        string `ini:"TokenKey"`     // 解析 token 的 aes key
	GoApiSignKey    string `ini:"GoApiSignKey"` //三方平台key
	Domains         string `ini:"Domains"`
	WebApiAccessKey string `ini:"WebApiAccessKey"` // callback访问key
}

// TybDbConf tydb
type TybDbConf struct {
	TSqlAddr       string `ini:"TSqlAddr"`
	TDbLogEnable   bool   `ini:"TDbLogEnable"`
	TDbMaxConnect  int    `ini:"TDbMaxConnect"`
	TDbIdleConnect int    `ini:"TDbIdleConnect"`
	TDbMaxLifeTime int    `ini:"TDbMaxLifeTime"`
}

// mysql 从库
type ReDbConf struct {
	RTSqlAddr       string `ini:"RTSqlAddr"`
	RTDbLogEnable   bool   `ini:"RTDbLogEnable"`
	RTDbMaxConnect  int    `ini:"RTDbMaxConnect"`
	RTDbIdleConnect int    `ini:"RTDbIdleConnect"`
	RTDbMaxLifeTime int    `ini:"RTDbMaxLifeTime"`
}

// RedisConf redis write
type RedisConf struct {
	RedisHost     string `ini:"RedisHost"`
	RedisAuth     string `ini:"RedisAuth"`
	RedisPoolSize int    `ini:"RedisPoolSize"`
}

// RedisConf redis read, 从库
type RRedisConf struct {
	RRedisHost     string `ini:"RRedisHost"`
	RRedisAuth     string `ini:"RRedisAuth"`
	RRedisPoolSize int    `ini:"RRedisPoolSize"`
}

// ConfigAddr 其它配置文件的路径
type ConfigAddr struct {
	IPConfAddr string `ini:"IPConfAddr"` // IP数据的配置文件
}

// LogConf 本地日志的配置
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
	LogType  string `ini:"LogType"`
}

type APNSConf struct {
	PEMFilePath string `ini:"PEMFilePath"`
	BundleId    string `ini:"BundleId"`
	KeyId       string `ini:"KeyId"`
	Issuer      string `ini:"KeyId"`
}

type S3Conf struct {
	AccessKeyID string `ini:"AccessKeyID"`
	SecretAccessKey    string `ini:"SecretAccessKey"`
	S3Region       string `ini:"S3Region"`
	S3Bucket      string `ini:"S3Bucket"`
}

func InitConfig(confPath *string) error {
	confInfo = new(Config)
	if err := ini.MapTo(confInfo, *confPath); err != nil {
		fmt.Println("MapTo", err)
		return err
	}
	return nil
}

func GetNatsConfig() NatsConfig {
	return confInfo.NatsConfig
}

// GetBaseConf 获取基础信息
func GetBaseConf() BaseConf {
	return confInfo.BaseConf
}

func GetSecretKey() string {
	return confInfo.SecretKey
}

// GetLogConf 获取日志的配置
func GetLogConf() LogConf {
	return confInfo.LogConf
}

// GetEnv 获取环境信息
func GetEnv() string {
	return confInfo.Env
}

func GetTyDbConf() TybDbConf {
	return confInfo.TybDbConf
}

func GetReDbConf() ReDbConf {
	return confInfo.ReDbConf
}

func GetRedisConf() RedisConf {
	return confInfo.RedisConf
}

func GetRRedisConf() RRedisConf {
	return confInfo.RRedisConf
}

// GetAuthFlag 获取鉴权开关
func GetAuthFlag() bool {
	return confInfo.AuthFlag
}

func GetAesSert() string {
	return confInfo.AesSert
}

func GetIPConfAddr() string {
	return confInfo.ConfigAddr.IPConfAddr
}

func GetSQDB() string {
	return confInfo.SqlDB
}

func GetTokenkey() string {
	return confInfo.BaseConf.TokenKey
}

func GetGoApiSignKey() string {
	return confInfo.BaseConf.GoApiSignKey
}

func GetFileUploadBase() string {
	return confInfo.FileStoragePath
}

func GetMaxUploadSize() int64 {
	return confInfo.MaxUploadSize
}

func GetUploadConf() UploadConf {
	return confInfo.UploadConf
}

func GetApnsConf() APNSConf {
	return confInfo.APNSConf
}

func GetS3Conf() S3Conf {
	return confInfo.S3Conf
}

