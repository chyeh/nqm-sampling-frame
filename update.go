package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

const (
	resetTargetAvailableSQL      = "UPDATE nqm_target SET tg_available=false;"
	resetTargetStatusSQL         = "UPDATE nqm_target SET tg_status=false;"
	updateNqmTargetAvailablePath = "var/sql/updateNqmTargetAvailable.sql"
	updateNqmTargetStatusPath    = "var/sql/updateNqmTargetStatus.sql"
)

func update() {
	if _, err := os.Stat(updateNqmTargetAvailablePath); os.IsNotExist(err) {
		log.Println("SQL file [", updateNqmTargetAvailablePath, "] doesn't exist")
		return
	}
	if _, err := os.Stat(updateNqmTargetStatusPath); os.IsNotExist(err) {
		log.Println("SQL file [", updateNqmTargetStatusPath, "] doesn't exist")
		return
	}

	if _, err := DB.Exec(resetTargetAvailableSQL); err != nil {
		log.Println(err)
	}
	sql, err := ioutil.ReadFile(updateNqmTargetAvailablePath)
	if err != nil {
		log.Println(err)
	}
	if _, err = DB.Exec(string(sql)); err != nil {
		log.Println(err)
	}

	if _, err := DB.Exec(resetTargetStatusSQL); err != nil {
		log.Println(err)
	}
	sql, err = ioutil.ReadFile(updateNqmTargetStatusPath)
	if err != nil {
		log.Println(err)
	}
	if _, err = DB.Exec(string(sql)); err != nil {
		log.Println(err)
	}
	log.Println("Updated MySQL")
}
