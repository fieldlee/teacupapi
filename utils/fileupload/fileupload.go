/*
* @ Author: Regee
* @ For:	分布式文件上传客户端
* @ Date:	2018年11月5日
* @ Des:	分布式文件上传客户端
 */
package fileupload

import (
	"io"
	"mime"

	yaboClient "teacupapi/utils/fileupload/seaweedfs"
)

// 文件上传接口
type YaoBoFileUploader interface {
	UploadFile(filename, mineType string, file io.Reader) (interface{}, error)
}

// 实例化一个上传客户端
func NewYaoBoFileUploader(fidURL, uploadURL string) YaoBoFileUploader {
	client := yaboClient.NewClient(fidURL, uploadURL)
	return client
}

//通过文件名称获取Mine type
func GetMineMimeType(fileExtension string) string {
	return mime.TypeByExtension(fileExtension)
}
