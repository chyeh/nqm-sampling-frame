package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	viper.SetConfigName("cfg")

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
