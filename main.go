package main

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	viper.SetConfigName("cfg")
	go pull(viper.GetString("fastbatServer"), viper.GetDuration("interval"))
	go update()
	select {}
}
