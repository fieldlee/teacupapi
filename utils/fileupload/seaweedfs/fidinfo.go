/*
* @ Author: Regee
* @ For:	seaweedfs 响应实体
* @ Date:	2018年11月5日
* @ Des:	seaweedfs 响应实体
 */
package seaweedfs

// seaweedfs 响应实体
type FidInfo struct {
	Fid       string `json:"fid"`       //Fid 可理解成服务器上面的文件名字
	URL       string `json:"url"`       // 获取fid 返回的url
	PublicURL string `json:"publicUrl"` // 获取fid 返回的 public url 现在服务器返回的是内网，直接在配置里面用外面ip去获取图片
	Count     int    `json:"count"`
}

// 上传响应实体
type UploadResp struct {
	FileName string `json:"name"`    // 文件名字
	FileURL  string `json:"fileURL"` // 文件完整路径
	Size     int64  `json:"size"`    // 文件大小
}
