package proxy

import (
	"log"
	"proxypool/database"
)

func init() {

	err := database.Engine.Sync2(&Proxy{})

	if err != nil {
		log.Fatal("sync database error", err)
	}
}

type Proxy struct {
	ID       int64  `xorm:"pk autoincr"`
	Ip       string `xorm:"NOT NULL"`
	Port     string `xorm:"NOT NULL"`
	Protocol string `xorm:"NOT NULL"`
	Latency  string

	Level    int
	UpdateAt string `xorm:"updated"`
	Refer    string `xorm:"NOT NULL"`
}

func (p * Proxy) Url()  string {
	return p.Protocol + "://" + p.Ip + ":" + p.Port
}

func Insert(proxy *Proxy) (err error) {
	c, err := database.Engine.Insert(proxy)
	log.Println("insert ", c)
	return
}