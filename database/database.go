package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

func SetupDatabase() error {
	var err error
	DbConn, err = sql.Open("mysql", "root:mysql-root@tcp(127.0.0.1:3306)/user_db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
	return nil
}
