package main

import (
	"io/ioutil"
	"log"
	"time"
)

func update(interval time.Duration) {
	ticker := time.NewTicker(time.Minute * interval)

	for {
		go func() {
			sql, err := ioutil.ReadFile("var/sql/updateNqmTargetAvailable.sql")
			if err != nil {
				log.Println(err)
			}
			if _, err = DB.Exec(string(sql)); err != nil {
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
