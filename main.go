package main

import (
	"fmt"
	"log"
	"proxypool/proxy"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	ch := make(chan proxy.Proxy)
	go proxy.A2uProvider(ch)
	go proxy.Data5uProvider(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
