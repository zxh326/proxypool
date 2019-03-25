package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"log"
	"proxypool/proxy"
	"sync"
	"time"
)

func main() {
	ch := make(chan *proxy.Proxy)

	checkTime := time.NewTicker(time.Minute * 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for  {
			fmt.Println("dddd")
			go  proxy.A2uProvider(ch)
			<- checkTime.C
		}
	}()

	go func() {
		for{
			p := <- ch
			go pings(p)
		}
	}()

	wg.Wait()

}

func pings(proxy *proxy.Proxy) {
	pinger, err := ping.NewPinger(proxy.Ip)
	pinger.Count = 3
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		proxy.Latency = stats.AvgRtt.String()
		log.Println(proxy)

	}
	pinger.Run()
}
