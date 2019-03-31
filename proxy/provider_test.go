package proxy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ch = make(chan *Proxy)

func TestHasProxy(t *testing.T) {
	proxy := Proxy{ID: 233, Ip: "123.233.233.233", Port: "8001", Protocol: "https", Level: 0, Refer: "A"}
	assert.Equal(t, HasProxyWithId(&proxy), false)
}

func TestInsertAndDelete(t *testing.T) {
	proxy := Proxy{ID: 2, Ip: "123.233.233.233", Port: "8001", Protocol: "https", Level: 0, Refer: "A"}

	Insert(&proxy)

	assert.Equal(t, proxy.Ip, GetOne(proxy.ID).Ip)

	Delete(&proxy)
}

func TestCount(t *testing.T) {
	fmt.Println(Count())
}
