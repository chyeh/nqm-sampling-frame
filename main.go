package main

import (
	"fmt"
	"time"
)

func updateMySQL() {
	for {
		fmt.Println("Update MySQL")
		time.Sleep(time.Second * 3)
	}
}

func pullFastBat() {
	for {
		fmt.Println("Pull from FastBat")
		time.Sleep(time.Second * 2)
	}
}

func main() {
	go pullFastBat()
	go updateMySQL()
	select {}
}
