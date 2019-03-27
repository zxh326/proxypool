package schedule

import (
	"github.com/parnurzeal/gorequest"
	"proxypool/common"
	"proxypool/proxy"
	"strconv"
	"time"
)

func ProxyValid(proxy proxy.Proxy) bool {
	var validUrl string
	if proxy.Protocol == "https"{
		validUrl = common.VerifyHttpsUrl
	}else {
		validUrl = common.VerifyUrl
	}

	begin := time.Now()
	res, _, err := gorequest.New().Proxy(proxy.Url()).Get(validUrl).Timeout(common.VaildTimeOut).End()
	defer res.Body.Close()
	if err != nil {
		return false
	}
	if res.StatusCode == 200 {
		proxy.Latency = strconv.Itoa(int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000)) + "ms"
		return true
	}
	return false
}
