package Dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(30 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)
}
