package main

import (
	"log"
	"os/exec"
	"time"
)

func pull(server string, interval time.Duration) {
	ticker := time.NewTicker(time.Minute * interval)

	for {
		go func() {
			log.Println("Pull from FastBat")
			exec.Command("./pull.sh", server).Run()
		}()
		<-ticker.C
	}

}
