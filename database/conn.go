package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", "./test.db")
	Engine.ShowSQL(true)
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	if err != nil {
		log.Fatal("init database err", err)
	}

	// 启动新线程定时激活sqlite
	go doPingEngine(time.Minute * 5)

}

func doPingEngine(sleep time.Duration) {
	for {
		time.Sleep(sleep)
		_ = Engine.Ping()
	}
}
