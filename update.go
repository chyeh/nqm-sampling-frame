package main

import (
	"io/ioutil"
	"log"
	"time"
)

const (
	resetTargetAvailableSQL = "UPDATE nqm_target SET tg_available=false;"
	resetTargetStatusSQL    = "UPDATE nqm_target SET tg_status=false;"
)

func update(interval time.Duration) {
	ticker := time.NewTicker(time.Minute * interval)

	for {
		go func() {
			if _, err := DB.Exec(resetTargetAvailableSQL); err != nil {
				log.Println(err)
			}
			sql, err := ioutil.ReadFile("var/sql/updateNqmTargetAvailable.sql")
			if err != nil {
				log.Println(err)
			}
			if _, err = DB.Exec(string(sql)); err != nil {
				log.Println(err)
			}

			if _, err := DB.Exec(resetTargetStatusSQL); err != nil {
				log.Println(err)
			}
			sql, err = ioutil.ReadFile("var/sql/updateNqmTargetStatus.sql")
			if err != nil {
				log.Println(err)
			}
			if _, err = DB.Exec(string(sql)); err != nil {
				log.Println(err)
			}
			log.Println("Updated MySQL")
		}()
		<-ticker.C
	}

}
