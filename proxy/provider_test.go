package proxy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
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
}

func TestInsertAndDelete(t *testing.T) {
	proxy := Proxy{Ip: "123.233.233.233", Port: "8000", Protocol: "https", Level: 0, Refer: "A"}

	Insert(&proxy)

	assert.Equal(t, proxy.Ip, GetOne(proxy.ID).Ip)

	Delete(&proxy)
}

func TestCheckServer(t *testing.T) {
	timeout := time.Duration(5 * time.Second)
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", "89.186.1.215:53281", timeout)
	fmt.Println("waist time :", time.Now().Sub(t1))
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
		return
	}
	fmt.Println("tcp server is ok")
}

//func TestOther(t *testing.T) {
//	GetSomething(1, 2, 3)
//
//}
