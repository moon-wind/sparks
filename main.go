package main

import (
	"demo2/utils"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func main() {
	f1()
}

func f1() {
	data := make(map[string]interface{})
	//data["account"] = "abcc"
	//data["passWord"] = "a123456"
	//data["captchaId"] = "vTaGqftKjdEObFBjeNzh"
	//data["captcha"] = "443872"

	//data["order_id"] = 48
	//data["answer"] = "1"

	client = &http.Client{}
	begin := time.Now()
	//url := "http://172.16.4.114:8888/pen/check"
	//url := "http://172.16.4.114:8888/awd_service/rfsc"
	url := "https://m.ccement.com/"
	//url := "http://172.16.4.114:9000/api/nine_user/login"
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := utils.CurlPost(client, url, data)
			fmt.Println(result)
		}()
	}
	wg.Wait()
	fmt.Printf("time consumed: %fs", time.Now().Sub(begin).Seconds())
}
