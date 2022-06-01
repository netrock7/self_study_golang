package main

import (
	"fmt"
	"net/http"
)

func main() { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url
	url := "http://gateway.solo.com/register/?id=999888777"
	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "xxx")
	reqest.Header.Add("X-Requested-With", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Body)
	defer response.Body.Close()
}
