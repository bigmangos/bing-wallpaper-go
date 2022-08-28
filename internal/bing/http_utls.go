package bing

import "github.com/go-resty/resty/v2"

func Get(url string) []byte {
	client := resty.New()
	resp, err := client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36").
		Get(url)
	if err != nil {
		panic(err)
	}
	return resp.Body()
}
