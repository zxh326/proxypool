package schedule

import (
	"github.com/parnurzeal/gorequest"
	"github.com/zxh326/proxypool/common"
	"github.com/zxh326/proxypool/proxy"
	"log"
	"time"
)

func init() {
	// 用于网络监测, 同时检测代理地址是否有效
	_, _, err := gorequest.New().Get(common.VerifyUrl).End()
	if err != nil {
		log.Fatal(err)
	}
}

func ProxyValid(proxy *proxy.Proxy) bool {
	var err []error
	var validUrl string
	var validCount, totalLatency = common.ValidCount, 0

	if proxy.Protocol == "https" {
		validUrl = common.VerifyHttpsUrl
	} else {
		validUrl = common.VerifyUrl
	}

	for i := 0; i < common.ValidCount; i++ {
		begin := time.Now()
		_, _, err = gorequest.New().Proxy(proxy.Url()).Get(validUrl).Timeout(common.ValidTimeOut).End()
		if err != nil {
			validCount--
			continue
		}
		totalLatency += int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000)
	}
	if validCount == 0 {
		if proxy.ID != 0 {
			InValidPool <- proxy
		}

		return false
	}
	log.Println("[proxy] valid one proxy success :", proxy.Url())
	proxy.Latency = totalLatency / validCount
	ValidPool <- proxy
	return true
}
