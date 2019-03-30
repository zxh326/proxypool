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
	Latency  int

	Level    int
	CreateAt string `xorm:"created"`
	Refer    string `xorm:"NOT NULL"`
}

func (p *Proxy) Url() string {
	return p.Protocol + "://" + p.Ip + ":" + p.Port
}

func Insert(proxy *Proxy) (err error) {
	if HasProxy(proxy) {
		return Update(proxy)
	}
	session := database.Engine.NewSession()
	defer session.Close()

	err = session.Begin()

	_, err = session.Insert(proxy)
	if err != nil {
		_ = session.Rollback()
		log.Println("[db] insert one proxy error: ", proxy, err)
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

func Update(proxy *Proxy) (err error) {
	session := database.Engine.NewSession()
	defer session.Close()

	err = session.Begin()
	_, err = session.ID(proxy.ID).Update(proxy)
	if err != nil {
		_ = session.Rollback()
		log.Println("[db] delete one proxy error: ", proxy)
		return
	}
	log.Println("[db] delete one invalid proxy: ", proxy)

	return session.Commit()
}

func HasProxy(proxy *Proxy) (has bool) {
	p := GetOne(proxy.ID)
	if p == nil {
		return false
	}
	return true
}

func Count() int64 {
	return int64(len(GetAll()))
}

func GetOne(id int64) *Proxy {
	tm := new(Proxy)
	c, _ := database.Engine.ID(id).Get(tm)
	if c {
		return tm
	}
	return nil
}

func GetAll(protocol ...string) (tm []*Proxy) {
	if len(protocol) == 0 {
		_ = database.Engine.Asc("Latency").Desc("create_at").Find(&tm)

	} else {
		_ = database.Engine.Asc("Latency").Desc("create_at").Where("protocol = ?", protocol[0]).Find(&tm)
	}
	return
}
