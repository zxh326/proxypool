package vaild

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"proxypool/proxy"
)

func ProxyTest(proxy proxy.Proxy){
	client := &http.Client{}

	var req *http.Request

	proxyFunc := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxy.Protocol+"://"+proxy.Ip+":"+proxy.Port) //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}
	transport := &http.Transport{Proxy: proxyFunc, ResponseHeaderTimeout:100 }
	client = &http.Client{Transport: transport}
	req, _ = http.NewRequest("get", "http://baidu.com", nil, )

	res, err := client.Do(req)

	defer res.Body.Close()
	if err != nil {
		log.Printf("%v, proxy vaild error", proxy)
		return
	}

	if err == nil{
		fmt.Println(res.StatusCode)
	}

}