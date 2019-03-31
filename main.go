package main

import (
	"github.com/zxh326/proxypool/api"
	"github.com/zxh326/proxypool/schedule"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	schedule.Start()
	api.Start()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
