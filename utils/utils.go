package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CurlPost 发送post请求
func CurlPost(client *http.Client, url string, data map[string]interface{}) string {
    bytesData, _ := json.Marshal(data)
    req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
    req.Header.Add("x-token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NDI2LCJVc2VybmFtZSI6ImFiY2MiLCJSZWFsTmFtZSI6ImFiYyIsIkNvbXBlSWQiOjEsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NDkzOTYwNjMsImlzcyI6ImNpbUVyQFpvMjIiLCJuYmYiOjE2NDg3OTAyNjN9.VuWVOu3-mp9lwQPaWporlZGhCJWVAD5GeWz9APHcELY")
    resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
    body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}