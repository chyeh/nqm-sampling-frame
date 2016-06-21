package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/chyeh/viper"
	"github.com/spf13/pflag"
	"github.com/toolkits/file"
)

func getFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	pflag.CommandLine.StringP("config", "c", "cfg.json", "configuration file")
	viper.BindPFlag("config", pflag.Lookup("config"))
}

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	cfgPath := viper.GetString("config")
	if !file.IsExist(cfgPath) {
		log.Fatalln("Configuration file [", cfgPath, "] doesn't exist")
	}

	viper.SetConfigName(getFileNameWithoutExtension(cfgPath))

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(dir)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	fastbatServer := viper.GetString("fastbatServer")
	interval := viper.GetDuration("interval")

	databaseInit()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		DB.Close()
		os.Exit(0)
	}()

	go pull(fastbatServer, interval)
	go update(interval)

	select {}
}
