/*
* @ Author: Regee
* @ For:	seaweedfs 客户端
* @ Date:	2018年11月5日
* @ Des:	seaweedfs 客户端
 */
package seaweedfs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/json-iterator/go"
)

var cjson = jsoniter.ConfigCompatibleWithStandardLibrary

// seaweed client
type Client struct {
	Fid       *FidInfo
	FidURL    string // http://ip:9333/dir/assign 配置文件配置好的请求url
	ManualURL string //手动配置上传url(不包含fid) 比如 http://ip:9342/fid 没有配置就用Fid里面返回的url
}

func NewClient(fidUrl, uploadUrl string) *Client {
	client := Client{}
	client.FidURL = fidUrl
	client.ManualURL = uploadUrl
	return &client
}

func (t *Client) getFid() error {

	if len(t.FidURL) < 0 {
		return errors.New("Client getFid invalid fidurl param")
	}
	resp, err := http.Get(t.FidURL)
	if err != nil {
		return errors.New("Client getFid url failed " + err.Error())
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Client->getFid->ReadAll request fild url failed " + err.Error())
	}
	cjson.Unmarshal(body, &t.Fid)
	return nil
}

func (t *Client) UploadFile(filename, mineType string, file io.Reader) (interface{}, error) {

	if t.Fid == nil || len(t.Fid.Fid) < 0 {
		err := t.getFid()
		if err != nil {
			return nil, err
		}
	}

	formData, contentType, err := t.MakeFormData(filename, mineType, file)
	if err != nil {
		return nil, err
	}

	url := t.Fid.PublicURL + "/" + t.Fid.Fid
	//if len(t.ManualURL) > 0 {
	//	if t.ManualURL[len(t.ManualURL)-1:] == "/" {
	//		url = t.ManualURL + t.Fid.Fid
	//	} else {
	//		url = t.ManualURL + "/" + t.Fid.Fid
	//	}
	//}
	fmt.Println(url)
	resp, err := t.upload(url, contentType, formData)
	if err != nil {
		return nil, err
	}

	resp.FileURL = url
	return *resp, nil
}

func (t *Client) upload(url string, contentType string, formData io.Reader) (r *UploadResp, err error) {
	resp, err := http.Post(url, contentType, formData)
	if err != nil {
		log.Println("Client->upload->" + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	uploadResp := new(UploadResp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Client->upload->ReadAll" + err.Error())
		return nil, err
	}
	cjson.Unmarshal(body, uploadResp)
	fmt.Println(resp, uploadResp)
	return uploadResp, nil
}

func (t *Client) MakeFormData(filename, mimeType string, context io.Reader) (formData io.Reader, contentType string, err error) {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	part, err := t.createFormFile(writer, "file", filename, mimeType)
	if err != nil {
		log.Println("Client->makeFormData->createFormFile->" + err.Error())
		return
	}
	_, err = io.Copy(part, context)
	if err != nil {
		log.Println("Client->makeFormData->Copy->" + err.Error())
		return
	}
	formData = buf
	contentType = writer.FormDataContentType()
	writer.Close()
	return
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func (t *Client) escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (t *Client) createFormFile(writer *multipart.Writer, fieldname, filename, mime string) (io.Writer, error) {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			t.escapeQuotes(fieldname), t.escapeQuotes(filename)))

	if len(mime) == 0 {
		mime = "application/octet-stream"
	}
	header.Set("Content-Type", mime)
	return writer.CreatePart(header)
}
