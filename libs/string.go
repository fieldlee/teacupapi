package libs

import (
	"strings"
	"teacupapi/config"
)

func TrimStaticDomain(url string) string {
	staticPicDomain := config.GetUploadConf().ImagesUrl
	//兼容第一个斜杠
	if string(staticPicDomain[len(staticPicDomain)-1]) != "/" {
		staticPicDomain = staticPicDomain + "/"
	}
	if strings.Contains(url, staticPicDomain) {
		return strings.ReplaceAll(url, staticPicDomain, "")
	} else {
		return url
	}
}

func TrimVideoStaticDomain(url string) string {
	staticPicDomain := config.GetUploadConf().VideoURL

	//兼容第一个斜杠
	if string(staticPicDomain[len(staticPicDomain)-1]) != "/" {
		staticPicDomain = staticPicDomain + "/"
	}

	if strings.Contains(url, staticPicDomain) {
		return strings.ReplaceAll(url, staticPicDomain, "")
	} else {
		return url
	}
}

func AppendStaticDomain(uri string) string {
	staticPicDomain := config.GetUploadConf().ImagesUrl
	if len(uri) == 0 ||
		len(staticPicDomain) == 0 {
		return ""
	}

	if strings.Contains(uri, staticPicDomain) {
		return uri
	}

	//去除域名最后一个斜杠
	if string(staticPicDomain[len(staticPicDomain)-1]) == "/" {
		staticPicDomain = staticPicDomain[0 : len(staticPicDomain)-1]
	}
	if string(uri[0]) == "/" {
		return staticPicDomain + uri
	} else {
		return staticPicDomain + "/" + uri
	}
}

func AppendVideoStaticDomain(uri string) string {
	staticPicDomain := config.GetUploadConf().VideoURL
	if len(uri) == 0 ||
		len(staticPicDomain) == 0 {
		return ""
	}

	if strings.Contains(uri, staticPicDomain) {
		return uri
	}

	//去除域名最后一个斜杠
	if string(staticPicDomain[len(staticPicDomain)-1]) == "/" {
		staticPicDomain = staticPicDomain[0 : len(staticPicDomain)-1]
	}
	if string(uri[0]) == "/" {
		return staticPicDomain + uri
	} else {
		return staticPicDomain + "/" + uri
	}
}
func AppendImageStaticDomain(uri string) string {
	staticPicDomain := config.GetUploadConf().ImagesUrl
	if len(uri) == 0 ||
		len(staticPicDomain) == 0 {
		return ""
	}

	if strings.Contains(uri, staticPicDomain) {
		return uri
	}

	//去除域名最后一个斜杠
	if string(staticPicDomain[len(staticPicDomain)-1]) == "/" {
		staticPicDomain = staticPicDomain[0 : len(staticPicDomain)-1]
	}
	if string(uri[0]) == "/" {
		return staticPicDomain + uri
	} else {
		return staticPicDomain + "/" + uri
	}
}

//反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
