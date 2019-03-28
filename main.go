package main

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/signal"
	"proxypool/schedule"
	"syscall"
)

func main() {
	schedule.Start()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
