package proxy

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"log"
	"proxypool/database"
)

func init()  {
	_  = database.Engine.Sync(&Proxy{})
}

type Proxy struct {
	Ip       string
	Port     string
	Protocol string
	Latency  string

	Level    int
	UpdateAt string
	refer    string
}

//func (p *Proxy) IsAvailable()  bool {
//	res, body, err := Request.Get("https://ip.cn/").Timeout(common.TimeOut).Proxy(p.Protocol+"://"+p.Ip+":"+p.Port).End()
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	fmt.Println(body, res.Body)
//	return true
//}

func (p *Proxy) ping(count int)  {
	pinger, err := ping.NewPinger(p.Ip)
	pinger.Count = count
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		p.Latency = stats.AvgRtt.String()
		log.Println(p)

	}
	pinger.Run()
}