package main

import (
	"time"

	"github.com/chyeh/viper"
)

func generate() {
	fastbatServer := viper.GetString("fastbatServer")
	interval := viper.GetDuration("interval")

	ticker := time.NewTicker(time.Minute * interval)
	for {
		pull(fastbatServer)
		update()

		<-ticker.C
	}
}
