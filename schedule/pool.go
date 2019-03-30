package schedule

import (
	"log"
	"proxypool/common"
	"proxypool/proxy"
	"time"
)

var (
	UnValidPool = make(chan *proxy.Proxy, 30)
	ValidPool   = make(chan *proxy.Proxy, 30)
)

func Start() {
	go Policy()
	go Valid()
	go Sync()
}

func Policy() {
	timeTicker := time.NewTicker(common.NextValidTime)

	for {
		// TODO policy
		CrawlerJob()
		<-timeTicker.C
	}

}

func Valid() {
	for ch := range UnValidPool {
		go ProxyValid(ch)
	}
}

func CrawlerJob() {
	providers := []func(ch chan<- *proxy.Proxy){
		proxy.Data5uProvider,
		proxy.A2uProvider,
	}

	log.Println("[crawler Job Begin]....")
	//var wg = sync.WaitGroup{}
	for _, providers := range providers {
		//wg.Add(1)
		//go func(f func(ch chan<- *proxy.Proxy)) {
		//	f(UnValidPool)
		//	wg.Done()
		//}(providers)
		providers(UnValidPool)
	}
	//wg.Wait()
	log.Println("[crawler Job End]....")
}

func Sync() {
	for p := range ValidPool {
		_ = proxy.Insert(p)
	}
}
