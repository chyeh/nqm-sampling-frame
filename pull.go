package main

import (
	"fmt"
	"time"
)

func pull() {
	for {
		fmt.Println("Pull from FastBat")
		time.Sleep(time.Second * 2)
	}
}
