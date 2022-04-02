package httpclient

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var (
	HttpClient *http.Client // 不过滤掉证书检查的 http client
	//SkipVerifyClient *http.Client // 过滤掉证书的  http client
)

// init HTTPClient  默认开启长链接(http 1.1之后) 开启http keepalive功能，也即是否重用连接，
func init() {
	HttpClient = createHTTPClient(true)
	//SkipVerifyClient = createHTTPClient(false)
}

const (
	MaxIdleConns        int = 100 // 连接池对所有host的最大链接数量，host也即dest-ip, 默认 100
	MaxIdleConnsPerHost int = 30  // 连接池对每个host的最大链接数量
	IdleConnTimeout     int = 90  // 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接（注意：该timeout起点是从每次空闲开始计时，若有交互则重置为0）,该参数通常设置为分钟级别
	RequestTimeout      int = 30  // 请求以及连接的超时时间
)

// createHTTPClient for connection re-use
func createHTTPClient(verifyFlag bool) *http.Client {
	if verifyFlag != true {
		return &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(RequestTimeout) * time.Second,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConns,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(RequestTimeout) * time.Second,
		}
	} else {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(RequestTimeout) * time.Second,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConns,
				MaxIdleConnsPerHost: MaxIdleConnsPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(RequestTimeout) * time.Second,
		}
	}
}

// GetHttpClient 获取 http client
func GetHttpClient(verifyFlag bool, connetTimeout time.Duration) *http.Client {
	return createNotMultiplexHTTPClient(verifyFlag, connetTimeout)
}

// createHTTPClient for connection re-use
func createNotMultiplexHTTPClient(verifyFlag bool, connetTimeout time.Duration) *http.Client {
	if connetTimeout <= 0 {
		connetTimeout = time.Duration(RequestTimeout) * time.Second
	}

	if !verifyFlag {
		return &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   connetTimeout,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,
				IdleConnTimeout: time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: connetTimeout,
		}
	} else {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   connetTimeout,
					KeepAlive: time.Duration(RequestTimeout) * time.Second,
				}).DialContext,

				IdleConnTimeout: time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: connetTimeout,
		}
	}
}
