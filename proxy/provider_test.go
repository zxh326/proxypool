package proxy

import (
	"fmt"
	"testing"
)

func TestData5uProvider(t *testing.T) {
	ch := make(chan *Proxy)
	go Data5uProvider(ch)
	proxy := <-ch
	fmt.Println(proxy)

}


func TestA2uProvider(t *testing.T) {
	ch := make(chan *Proxy)
	go A2uProvider(ch)
	proxy := <-ch
	fmt.Println(proxy)

}
