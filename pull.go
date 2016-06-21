package main

import (
	"log"
	"os/exec"
)

func pull(server string) {
	log.Println("Run /pull.sh for pulling sampling frame from FastBat")
	exec.Command("./pull.sh", server).Run()
}
