package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/go_test")
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println("Database connection successful.....")
	return db
}
