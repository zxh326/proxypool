package main

import (
	"os"
	"os/signal"
	"proxypool/api"
	"proxypool/schedule"
	"syscall"
)

func main() {
	schedule.Start()
	api.Start()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
