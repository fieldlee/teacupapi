package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"teacupapi/db/mdata"
)

// POSTHTTP post 请求返回 json
func POSTHTTP(url string, dataBytes []byte, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(dataBytes))
	if err != nil {
		return []byte(""), err
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

// GETHTTP get 请求返回 json
func GETHTTP(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte(""), err
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func POSTJson(path string, postData map[string]interface{}, header map[string]string) ([]byte, error) {

	jsonStr, err := mdata.Cjson.Marshal(postData)
	if err != nil {
		return []byte(""), err
	}

	payload := strings.NewReader(string(jsonStr))
	req, _ := http.NewRequest("POST", path, payload)
	req.Header.Add("content-type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}
	/*
	   req.Header.Add("isupdate", isupdate)
	   req.Header.Add("Api-Key", apiKey)
	   req.Header.Add("accept", "application/json")
	*/

	resp, err := HttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}
