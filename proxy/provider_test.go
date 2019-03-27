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
	fmt.Println(proxy)

	Insert(proxy)
}

func TestInsert(t *testing.T) {
	proxy := Proxy{Ip:"123.233.233.233", Port:"8000", Protocol:"https", Level:0, }

	Insert(&proxy)
}