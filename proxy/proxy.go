package proxy

import (
	"fmt"
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

func (p *Proxy) Url() string {
	return p.Protocol + "://" + p.Ip + ":" + p.Port
}

func Insert(proxy *Proxy) (err error) {
	session := database.Engine.NewSession()
	defer session.Close()

	err = session.Begin()

	_, err = session.Insert(proxy)
	if err != nil {
		_ = session.Rollback()
		log.Println("[db] insert one proxy error: ", proxy)
		return
	}
	return session.Commit()
}

func Delete(proxy *Proxy) (err error) {
	session := database.Engine.NewSession()
	defer session.Close()

	err = session.Begin()
	_, err = session.ID(proxy.ID).Delete(proxy)
	if err != nil {
		_ = session.Rollback()
		log.Println("[db] delete one proxy error: ", proxy)
		return
	}
	log.Println("[db] delete one invalid proxy: ", proxy)

	return session.Commit()
}

func Count() int64 {
	return int64(len(GetAll()))
}


func GetOne(id int64) *Proxy {
	tm := new(Proxy)
	c, _ := database.Engine.ID(id).Get(tm)
	fmt.Println(c)
	return tm
}

func GetAll(protocol ...string) (tm []*Proxy) {
	if len(protocol) == 0{
		_ = database.Engine.Desc("update_at").Asc("Latency").Find(&tm)

	}else{
		_ = database.Engine.Desc("update_at").Asc("Latency").Where("protocol = ?", protocol[0]).Find(&tm)
	}
	return
}
