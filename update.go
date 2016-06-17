package main

import (
	"fmt"
	"time"
)

func update() {
	for {
		fmt.Println("Update MySQL")
		time.Sleep(time.Second * 3)
	}
}
