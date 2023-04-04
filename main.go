package main

import (
	"EnployeeService/config"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.ReadCfg()
	cxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := NewApp(cfg, cxt)
	app.InitServices()
	app.Start()

	doneSignal := make(chan os.Signal, 1)
	signal.Notify(doneSignal, syscall.SIGTERM, os.Interrupt)
	<-doneSignal
}
