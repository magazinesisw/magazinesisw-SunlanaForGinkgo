package Utils

import (
	"github.com/go-resty/resty/v2"
	"log"
)

// SetRequest 封装请求参数
func SetRequest(client *resty.Client, url string, method string, postData interface{}) (string, int, error) {

	// 创建请求
	req := client.R().
		SetHeader("Content-Type", "application/json"). // 设置请求头
		SetBody(postData)

	// 发送请求
	var resp *resty.Response
	var err error

	switch method {
	case "POST":
		resp, err = req.Post(url)
	case "GET":
		resp, err = req.Get(url)
	case "PUT":
		resp, err = req.Put(url)

	default:
		log.Println("method not support")

	}
	if err != nil {
		log.Println(err)
		return "", 0, err
	}

	return resp.String(), resp.StatusCode(), nil

}
