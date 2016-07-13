package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cepave/open-falcon-backend/common/logruslog"
	"github.com/Cepave/open-falcon-backend/common/vipercfg"
	"github.com/spf13/pflag"
)

func main() {
	vipercfg.Parse()
	vipercfg.Bind()

	if vipercfg.Config().GetBool("version") {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if vipercfg.Config().GetBool("help") {
		pflag.Usage()
		os.Exit(0)
	}

	vipercfg.Load()
	logruslog.Init()
	dbInit()

	go generate()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		DB.Close()
		os.Exit(0)
	}()

	select {}
}
