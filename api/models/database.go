package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func Connect() {
	Db, err = sql.Open("mysql", "root:pere@123@/go")
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {
	err = Db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
