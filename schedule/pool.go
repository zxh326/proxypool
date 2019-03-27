package schedule

import (
	"log"
	"proxypool/common"
	"proxypool/proxy"
	"sync"
	"time"
)

var (
	UnValidPool = make(chan *proxy.Proxy)
	ValidPool = make(chan *proxy.Proxy)
)

func Start()  {
	go Policy()

}

func CrawlerJob()  {
	providers := []func(ch chan<- *proxy.Proxy)  {
		proxy.Data5uProvider,
		proxy.A2uProvider,
	}

	log.Println("crawler Job Begin....")
	var wg = sync.WaitGroup{}
	for _, providers := range providers {
		wg.Add(1)
		go func() {
			providers(UnValidPool)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("crawler Job End....")

	//for   {
	//
	//}
}

func Policy(){
	timeTicker := time.NewTicker(common.NextVaildTime)

	for {
		// TODO policy
		CrawlerJob()
		<-timeTicker.C
	}


}