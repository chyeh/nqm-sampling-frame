package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chyeh/viper"
	"github.com/spf13/pflag"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	pflag.CommandLine.StringP("config", "c", "cfg.json", "configuration file")
	viper.BindPFlag("config", pflag.Lookup("config"))
	pflag.CommandLine.BoolP("version", "v", false, "show version")
	viper.BindPFlag("version", pflag.Lookup("version"))
}

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if viper.GetBool("version") {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	loadConfigFile()

	fastbatServer := viper.GetString("fastbatServer")
	interval := viper.GetDuration("interval")

	dbInit()

	go pull(fastbatServer, interval)
	go update(interval)

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
