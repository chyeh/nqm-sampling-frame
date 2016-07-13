package main

import (
	"database/sql"
	"fmt"

	"github.com/Cepave/open-falcon-backend/common/vipercfg"
	log "github.com/Sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func dbConnInit(dsn string) (err error) {
	if DB, err = sql.Open("mysql", dsn); err != nil {
		return fmt.Errorf("Open DB error: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("Ping DB error: %v", err)
	}

	return
}

func dbInit() {
	err := dbConnInit(vipercfg.Config().GetString("database"))

	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxIdleConns(100)
}
