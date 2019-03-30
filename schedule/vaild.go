package schedule

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"proxypool/common"
	"proxypool/proxy"
	"strconv"
	"time"
)

func ProxyValid(proxy *proxy.Proxy) bool {
	var validUrl string
	if proxy.Protocol == "https" {
		validUrl = common.VerifyHttpsUrl
	} else {
		validUrl = common.VerifyUrl
	}

	begin := time.Now()
	res, _, err := gorequest.New().Proxy(proxy.Url()).Get(validUrl).Timeout(common.ValidTimeOut).End()
	if err != nil {
		log.Println("[proxy] ", proxy, " valid failed")
		return false
	}
	if res.StatusCode == 200 {
		proxy.Latency = strconv.Itoa(int(time.Now().Sub(begin).Nanoseconds()/1000/1000)) + "ms"
		ValidPool <- proxy
		return true
	}
	defer res.Body.Close()
	return false
}
