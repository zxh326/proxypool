package database

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var Engine *xorm.Engine

func init()  {
	Engine, err := xorm.NewEngine("sqlite3", "./test.db")

	if err != nil{
		log.Fatal("init database err", err)
	}

	go doPingEngine(time.Minute * 5)

}

func doPingEngine(sleep time.Duration)  {
	for {
		_ = Engine.Ping()
		time.Sleep(sleep)
	}
}