package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"teacupapi/config"
	"teacupapi/controllers/base"
	"teacupapi/libs"
	glog "teacupapi/logs"
	"teacupapi/utils"
	"time"
)

func VideoUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		glog.Errorf("获取文件错误，err:%+v", err)
		base.WebResp(c, 200, utils.ErrInternal, nil, "存储文件错误")
		return
	}
	if file.Size == 0 {
		base.WebResp(c, 200, utils.ErrInternal, nil, "文件大小为空")
		return
	}

	if file.Size > 1024*1024*1024 {
		base.WebResp(c, 200, utils.ErrInternal, nil, "视频文件不许超过1GB")
		return
	}

	destPath := config.GetUploadConf().VideoStoragePath
	fileSuffix := path.Ext(file.Filename)
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	storagePath := filepath.FromSlash(destPath + "/" + filename)
	err = c.SaveUploadedFile(file, storagePath)
	if err != nil {
		glog.Errorf("视频上传保存文件，err:%+v", err)
		base.WebResp(c, 200, utils.ErrInternal, nil, "存储文件错误")
		return
	}
	base.WebResp(c, http.StatusOK, utils.StatusOK, libs.AppendVideoStaticDomain(filename), utils.Success)
}

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		glog.Errorf("获取文件错误，err:%+v", err)
		base.WebResp(c, 200, utils.ErrInternal, nil, "存储文件错误")
		return
	}
	if file.Size == 0 {
		base.WebResp(c, 200, utils.ErrInternal, nil, "文件大小为空")
		return
	}

	if file.Size > 1024*1024*3 {
		base.WebResp(c, 200, utils.ErrInternal, nil, "图片大小不允许超过3M")
		return
	}

	destPath := config.GetUploadConf().FileStoragePath
	fileSuffix := path.Ext(file.Filename)

	if !VerifyImageFormat(fileSuffix) {
		base.WebResp(c, 200, utils.ErrInternal, nil, "上传图片只能为gif，png，jpg格式")
		return
	}

	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	storagePath := filepath.FromSlash(destPath + "/" + filename)
	err = c.SaveUploadedFile(file, storagePath)
	if err != nil {
		glog.Errorf("图片上传保存文件，err:%+v", err)
		base.WebResp(c, 200, utils.ErrInternal, nil, "存储文件错误")
		return
	}
	base.WebResp(c, http.StatusOK, utils.StatusOK, libs.AppendImageStaticDomain(filename), utils.Success)
}

func ImageUpload(c *gin.Context) {
	code, msg := seaweedUpload(c)
	if code == utils.StatusOK {
		base.WebResp(c, 200, utils.StatusOK, msg, utils.Success)
	} else {
		base.WebResp(c, 200, code, nil, msg)
	}
	return
}

// 上传文件,图片参数name为img
func seaweedUpload(c *gin.Context) (error int, message string) {
	var fidURL = config.GetUploadConf().FileStoragePath
	var filename = utils.GetRandomString(10) + ".jpg"

	file, err := c.FormFile("file")
	if err != nil {
		return utils.ErrInvalidParams, "参数错误"
	}
	fileSuffix := path.Ext(file.Filename)
	if !VerifyImageFormat(fileSuffix) {
		return utils.ErrInvalidParams, "上传图片只能为gif，png，jpg格式"
	}
	fileData, err := file.Open()
	if err != nil {
		return utils.ErrInvalidParams, ""
	}
	defer func() { _ = fileData.Close() }()

	saveFile, err := os.Create(fmt.Sprintf("%s/%s", fidURL, filename))
	if err != nil {
		glog.Errorf("UploadHandler: 创建文件失败！")
		return utils.ErrInvalidParams, err.Error()
	}
	defer saveFile.Close()

	_, errCopy := io.Copy(saveFile, fileData)
	if errCopy != nil {
		glog.Errorf("UploadHandler: 文件复制失败！ -> {%s}", err)
		return utils.ErrInvalidParams, err.Error()
	}
	return utils.StatusOK, fmt.Sprintf("/%s", filename)
}

func VerifyImageFormat(fileSuffix string) bool {
	switch fileSuffix {
	case ".gif", ".png", ".jpg", ".jpeg":
		return true
	default:
		return false
	}
}
