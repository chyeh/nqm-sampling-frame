package main

import (
	"time"

	"github.com/Cepave/open-falcon-backend/common/vipercfg"
)

func generate() {
	fastbatServer := vipercfg.Config().GetString("fastbatServer")
	interval := vipercfg.Config().GetDuration("interval")

	ticker := time.NewTicker(time.Minute * interval)
	for {
		pull(fastbatServer)
		update()

		<-ticker.C
	}
}
