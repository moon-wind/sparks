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
	req.Header.Add("x-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NTcsIlVzZXJuYW1lIjoiMTg5NTI5Mzg4MjgiLCJSZWFsTmFtZSI6IuW4uOiKsyIsIkNvbXBlSWQiOjYzLCJCdWZmZXJUaW1lIjo4NjQwMCwiZXhwIjoxNjcxNjEwNDMzLCJpc3MiOiJjaW1FckBabzIyIiwibmJmIjoxNjcxMDA0NjMzfQ.bzck_lRbYrgl9Sn1JFOSDVIi-tS8Arbu6NUS1BHCKjo")
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
