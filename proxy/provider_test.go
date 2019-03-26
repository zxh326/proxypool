package proxy

import (
	"fmt"
	"testing"
)

var ch = make(chan *Proxy)

func TestData5uProvider(t *testing.T) {
	go Data5uProvider(ch)
	proxy := <-ch
	fmt.Println(proxy)
}


func TestA2uProvider(t *testing.T) {
	go A2uProvider(ch)
	proxy := <-ch
	proxy.ping(3)
}
