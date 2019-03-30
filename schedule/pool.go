package schedule

import (
	"log"
	"proxypool/common"
	"proxypool/proxy"
	"sync"
	"time"
)

var (
	UnValidPool = make(chan *proxy.Proxy, 30)
	ValidPool   = make(chan *proxy.Proxy, 30)
	InValidPool = make(chan *proxy.Proxy, 30)
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
		var wg = sync.WaitGroup{}
		for _, p := range proxy.GetAll() {
			wg.Add(1)
			go func(p *proxy.Proxy) {
				ProxyValid(p)
				wg.Done()
			}(p)
		}
		wg.Wait()

		if proxy.Count() <= common.MinPoolNum {
			CrawlerJob()
		}
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
		proxy.LiuLiuProvider,
	}

	log.Println("[crawler Job Begin]....")
	for _, providers := range providers {
		providers(UnValidPool)
	}
	log.Println("[crawler Job End]....")
}

func Sync() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for p := range ValidPool {
			_ = proxy.Insert(p)
		}
	}()

	go func() {
		for p := range InValidPool {
			_ = proxy.Delete(p)
		}
	}()
	wg.Wait()
}
