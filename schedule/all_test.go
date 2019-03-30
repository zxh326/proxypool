package schedule

import (
	"fmt"
	"proxypool/proxy"
	"testing"
)

func TestProxyValid(t *testing.T) {
	proxy_ := proxy.Proxy{Ip: "123.233.233.233", Port: "8000", Protocol: "https", Level: 0, Refer: "A"}

	fmt.Println(ProxyValid(&proxy_))

}
