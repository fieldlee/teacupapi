# 文件上传测试Demo 

创建一个main.go 具体使用方法如下,网络请求 从http.Request方法获取FormFile("file") 获取File

如果是用过配置配置fidURL,uploadURL 在regge目录下config.json 有配置示例
"fidURL":"http://47.244.62.69:9333/dir/assign",  
"fileUploadBase":"http://47.244.62.69:9342"
获取配置值
 fidURL:= yabo.Conf["fidURL"]
 fileUploadBase:= yabo.Conf["fileUploadBase"]
    
```
var (
	filename  = "test.jpg"
	fidURL    = "http://47.244.62.69:9333/dir/assign"
	uploadURL = "http://47.244.62.69:9342"
)

func main() {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	upload := yaboUpload.NewYaoBoFileUploader(fidURL, uploadURL)
	resp, err := upload.UploadFile(filename, "image/jpg", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
```