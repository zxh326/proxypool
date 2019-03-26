package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
