package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chyeh/viper"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func databaseInit() {
	err := dbInit(viper.GetString("database"))

	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxIdleConns(100)
}

func dbInit(dsn string) (err error) {
	if DB, err = sql.Open("mysql", dsn); err != nil {
		return fmt.Errorf("Open DB error: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("Ping DB error: %v", err)
	}

	return
}

// Convenient IoC for transaction processing
func inTx(txCallback func(tx *sql.Tx) error) (err error) {
	var tx *sql.Tx

	if tx, err = DB.Begin(); err != nil {
		return
	}

	/**
	 * The transaction result by whether or not the callback has error
	 */
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	// :~)

	err = txCallback(tx)

	return
}
