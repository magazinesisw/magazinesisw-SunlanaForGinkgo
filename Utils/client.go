package Utils

import "github.com/go-resty/resty/v2"

func SetClient() *resty.Client {

	var client *resty.Client
	client = resty.New()

	return client

}
